package usecase

import (
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

var _ blacklist.IAddBlacklist = (*UsecaseAddBlacklist)(nil)

type UsecaseAddBlacklist struct {
	blacklist_repository repositoty.IBlackListRepository
}

func NewAddBlacklist(blacklist_repository repositoty.IBlackListRepository) *UsecaseAddBlacklist {
	return &UsecaseAddBlacklist{blacklist_repository: blacklist_repository}
}

func (c *UsecaseAddBlacklist) Execute(userIdentifier int, eventId, reason, document, scope string, blocked_until *time.Time) (*entity.BlackList, error) {
	var blacklistEmpty = entity.BlackList{}

	blocked_type := entity.TEMPORARY
	if blocked_until == nil {
		blocked_type = entity.PERMANENT
	}

	blacklist := entity.NewBlackList(eventId, reason, document, scope, blocked_type, userIdentifier, blocked_until)
	err := blacklist.IsValid()
	if err != nil {
		return &blacklistEmpty, err
	}
	err = c.blacklist_repository.Add(blacklist)
	if err != nil {
		return &blacklistEmpty, err
	}
	return blacklist, nil
}
