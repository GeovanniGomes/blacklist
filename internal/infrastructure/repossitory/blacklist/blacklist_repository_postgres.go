package blacklist

import (
	"fmt"
	"log"
	"sync"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

var _ repositoty.BlackListRepositoryInterface = (*BlackListRepositoryPostgres)(nil)

type BlackListRepositoryPostgres struct {
	mutex       sync.Mutex
	persistence contracts.DatabaseRelationalInterface
}

func NewBlackListRepositoryPostgres(persistence contracts.DatabaseRelationalInterface) *BlackListRepositoryPostgres {
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
		log.Fatalf("Error peersiste data blacklist: %v", err)
		return err
	}
	return nil
}

func (b *BlackListRepositoryPostgres) Check(userIdentifier int, evendId string) (bool, string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	rows, err := b.persistence.SelectQuery("SELECT reason FROM blacklist WHERE user_identifier = $1 and event_id = $2 and is_active = $3", userIdentifier, evendId, true)

	if err != nil {
		log.Fatalf("Error querying blacklist: %v", err)
		panic(fmt.Sprintf("Error check in database: %v, %v", userIdentifier, evendId))
	}
	defer rows.Close()

	var reason string
	if rows.Next() {
		if err := rows.Scan(&reason); err != nil {
			return false, ""
		}
		return false, reason
	}
	return true, ""

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
