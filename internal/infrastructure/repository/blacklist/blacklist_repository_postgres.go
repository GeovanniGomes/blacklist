package blacklist

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

var _ repositoty.IBlackListRepository = (*BlackListRepositoryPostgres)(nil)

type BlackListRepositoryPostgres struct {
	mutex       sync.Mutex
	persistence contracts.IDatabaseRelational
}

func (b *BlackListRepositoryPostgres) FetchBlacklistEntries(startDate, endDate time.Time) ([]entity.BlackList, error) {
	factory := entity.FactoryEntity{}
	query := `SELECT * FROM blacklist WHERE created_at BETWEEN $1 AND $2 and is_active = $3`
	rows, err := b.persistence.SelectQuery(query, startDate, endDate, true)
	if err != nil {
		log.Printf("error fetch blacklist: %v", err)
		return nil, errors.New("unable to obtain blacklist data")
	}
	defer rows.Close()

	var blacklists []entity.BlackList
	for rows.Next() {
		var (
			id             *string
			eventId        *string
			createdAt      time.Time
			reason         string
			document       string
			scope          string
			userIdentifier int
			blockedUntil   *time.Time
			blockedType    string
			isActive       bool
		)
		if err := rows.Scan(&id, &eventId, &createdAt, &reason, &document, &scope, &userIdentifier, &blockedUntil, &blockedType, &isActive); err != nil {
			return nil, fmt.Errorf("erro ao escanear linha: %v", err)
		}
		blacklist, err := factory.FactoryNewBlacklist(eventId, reason, document, scope, blockedType, userIdentifier, isActive, blockedUntil, &createdAt, id)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar entidade BlackList: %v", err)
		}
		blacklists = append(blacklists, *blacklist)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre os resultados: %v", err)
	}

	return blacklists, nil
}

func NewBlackListRepositoryPostgres(persistence contracts.IDatabaseRelational) *BlackListRepositoryPostgres {
	return &BlackListRepositoryPostgres{persistence: persistence}
}

func (b *BlackListRepositoryPostgres) AddBlacklist(blacklist *entity.BlackList) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	err := b.persistence.InsertData(
		"blacklist",
		[]string{"id", "user_identifier", "event_id", "scope", "reason", "document", "blocked_until", "blocked_type", "is_active"},
		[]interface{}{
			blacklist.GetId(),
			blacklist.GetUserIdentifier(),
			blacklist.GetEventId(),
			blacklist.GetScope(),
			blacklist.GetReason(),
			blacklist.GetDocument(),
			blacklist.GetBlockedUntil(),
			blacklist.GetBlockedType(),
			blacklist.GetIsActive(),
		},
	)
	if err != nil {
		log.Printf("Error peersiste data blacklist: %v", err)
		return err
	}
	return nil
}

func (b *BlackListRepositoryPostgres) CheckBlacklist(userIdentifier int, eventId *string) (*entity.BlackList, error) {
	factory := entity.FactoryEntity{}
	b.mutex.Lock()
	defer b.mutex.Unlock()

	query := "SELECT *FROM blacklist WHERE user_identifier = $1 and is_active = $2"
	args := []interface{}{userIdentifier, true}

	if eventId != nil {
		query += " and event_id = $3"
		args = append(args, eventId)
	}
	query += " ORDER BY created_at DESC LIMIT 1"
	rows, err := b.persistence.SelectQuery(query, args...)

	if err != nil {
		log.Fatalf("Error querying blacklist: %v", err)
		return &entity.BlackList{}, errors.New("unable to complete blacklist check")
	}
	defer rows.Close()
	var blacklists []entity.BlackList
	for rows.Next() {
		var (
			id                 *string
			scanEventId        *string
			createdAt          time.Time
			reason             string
			document           string
			scope              string
			scanUserIdentifier int
			blockedUntil       *time.Time
			blockedType        string
			isActive           bool
		)
		if err := rows.Scan(&id, &scanEventId, &createdAt, &reason, &document, &scope, &scanUserIdentifier, &blockedUntil, &blockedType, &isActive); err != nil {
			return &entity.BlackList{}, fmt.Errorf("erro ao escanear linha: %v", err)
		}

		blacklist, err := factory.FactoryNewBlacklist(eventId, reason, document, scope, blockedType, userIdentifier, isActive, blockedUntil, &createdAt, id)

		if err != nil {
			return nil, fmt.Errorf("erro ao criar entidade BlackList: %v", err)
		}
		blacklists = append(blacklists, *blacklist)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre os resultados: %v", err)
	}
	if len(blacklists) == 0 {
		return &entity.BlackList{}, nil
	}
	return &blacklists[0], nil
}

func (b *BlackListRepositoryPostgres) RemoveBlacklist(userIdentifier int, eventId *string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	defaultQuery := "user_identifier = $2"

	if eventId !=nil{
		defaultQuery += " AND event_id = $3"
		return b.persistence.UpdateData(
			"blacklist",
			[]string{"is_active"},
			[]interface{}{false},
			defaultQuery,
			userIdentifier, eventId,
		)
	}
	return b.persistence.UpdateData(
		"blacklist",
		[]string{"is_active"},
		[]interface{}{false},
		defaultQuery,
		userIdentifier,
	)
}

func (b *BlackListRepositoryPostgres) AddEvent(event entity.Event) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	err := b.persistence.InsertData(
		"events",
		[]string{"id", "title", "description", "date", "category", "is_active", "created_at"},
		[]interface{}{
			event.GetId(),
			event.GetTitle(),
			event.GetDescription(),
			event.GetDate(),
			event.GetCategory().GetName(),
			event.GetIsActive(),
			event.GetCreatedAt(),
		},
	)
	if err != nil {
		log.Printf("Error peersiste data event: %v", err)
		return err
	}
	return nil
}

func (b *BlackListRepositoryPostgres) GetEvent(id string) (*entity.Event, error) {
	factory := entity.FactoryEntity{}
	b.mutex.Lock()
	defer b.mutex.Unlock()

	query := "SELECT * FROM events WHERE id = $1 and is_active = $2"
	args := []interface{}{id, true}

	rows, err := b.persistence.SelectQuery(query, args...)
	if err != nil {
		log.Printf("Error querying events: %v", err)
		return nil, errors.New("unable to complete event get")
	}
	defer rows.Close()

	var (
		idEvent     string
		title       string
		description string
		date        time.Time
		category    string
		isActive    bool
		createdAt   time.Time
	)

	if rows.Next() {
		if err := rows.Scan(&idEvent, &title, &description, &date, &category, &isActive, &createdAt); err != nil {
			return nil, fmt.Errorf("erro ao escanear linha: %v", err)
		}

		event, err := factory.FactoryNewEvent(&idEvent, title, description, date, &createdAt, category, isActive)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar entidade BlackList: %v", err)
		}
		return event, nil
	}

	return nil, nil
}

func (b *BlackListRepositoryPostgres) RemoveEvent(id string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.persistence.UpdateData(
		"events",
		[]string{"is_active"},
		[]interface{}{false},
		"id = $2",
		id,
	)
}
