package usecase

import (
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

var _ blacklist.ICheckBlacklist = (*UsecaseCheckBlacklist)(nil)

type UsecaseCheckBlacklist struct {
	blacklist_repository repository.IBlackListRepository
}

func NewCheckBlacklist(
	blacklist_repository repository.IBlackListRepository) *UsecaseCheckBlacklist {
	return &UsecaseCheckBlacklist{blacklist_repository: blacklist_repository}
}

func (c *UsecaseCheckBlacklist) Execute(userIdentifier int, eventId *string) (string, error) {
	blaclistEntity, err := c.blacklist_repository.CheckBlacklist(userIdentifier, eventId)

	if err != nil {

	}
	if blaclistEntity != nil {
		if blaclistEntity.GetScope() == entity.GLOBAL {
			return blaclistEntity.GetReason(), nil
		}

		if blaclistEntity.GetEventId() == eventId {
			return blaclistEntity.GetReason(), nil
		}
	}
	return "", nil
}
