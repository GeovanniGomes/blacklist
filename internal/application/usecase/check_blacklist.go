package usecase

import (
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
)

var _ blacklist.ICheckBlacklist = (*UsecaseCheckBlacklist)(nil)

type UsecaseCheckBlacklist struct {
	blacklist_repository repository.IBlackListRepository
}

func NewCheckBlacklist(
	blacklist_repository repository.IBlackListRepository) *UsecaseCheckBlacklist {
	return &UsecaseCheckBlacklist{blacklist_repository: blacklist_repository}
}

func (c *UsecaseCheckBlacklist) Execute(userIdentifier int, eventId string) (string, error) {
	return c.blacklist_repository.Check(userIdentifier, eventId)
}
