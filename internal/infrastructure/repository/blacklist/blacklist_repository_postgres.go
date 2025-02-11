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
			eventId        string
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

func (b *BlackListRepositoryPostgres) Add(blacklist *entity.BlackList) error {
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

func (b *BlackListRepositoryPostgres) Check(userIdentifier int, evendId string) (string, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	rows, err := b.persistence.SelectQuery("SELECT reason FROM blacklist WHERE user_identifier = $1 and event_id = $2 and is_active = $3", userIdentifier, evendId, true)

	if err != nil {
		log.Fatalf("Error querying blacklist: %v", err)
		return "", errors.New("unable to complete blacklist check")
	}
	defer rows.Close()

	var reason string
	if rows.Next() {
		if err := rows.Scan(&reason); err != nil {
			return "", errors.New("unable to complete blacklist check")
		}
		return reason, nil
	}
	return "", nil

}

func (b *BlackListRepositoryPostgres) Remove(userIdentifier int, eventId string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.persistence.UpdateData(
		"blacklist",
		[]string{"is_active"},
		[]interface{}{false},
		"user_identifier = $2 AND event_id = $3",
		userIdentifier, eventId,
	)
}
