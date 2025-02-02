package blacklist_repository

import (
	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence/contracts"
	"log"
)

var _ repositoty.BlackListRepositoryInterface = (*BlackListRepositoryPostgres)(nil)

type BlackListRepositoryPostgres struct {
	persistence contracts.DatabaseRelationalInterface
}
func NewBlackListRepositoryPostgres(persistence contracts.DatabaseRelationalInterface) *BlackListRepositoryPostgres {
	return &BlackListRepositoryPostgres{persistence: persistence}
}

func (b *BlackListRepositoryPostgres) Add(blacklist *entity.BlackList) error {
	err := b.persistence.InsertData(
		"blacklist",
		[]string{"id", "user_identifier", "event_id", "scope", "document", "blocked_until", "blocked_type"},
		[]interface{}{
			blacklist.GetId(),
			blacklist.GetUserIdentifier(),
			blacklist.GetEventId(),
			blacklist.GetScope(),
			blacklist.GetDocument(),
			blacklist.GetBlockedUntil(),
			blacklist.GetBlockedType(),
		},
	)
	if err != nil {
			log.Fatalf("Error peersiste data blacklist: %v", err)
		return err
	}
	return nil
}

// Check implements repositoty.BlackListRepositoryInterface.
func (b *BlackListRepositoryPostgres) Check(userIndentifier int, evendId string) (bool, string) {
	panic("unimplemented")
}

// Remove implements repositoty.BlackListRepositoryInterface.
func (b *BlackListRepositoryPostgres) Remove(userIndentifier int, eventId string) error {
	panic("unimplemented")
}


