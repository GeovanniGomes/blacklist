package usecase

import (
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
)

var _ blacklist.CheckBlacklistInterface = (*UsecaseCheckBlacklist)(nil)

type UsecaseCheckBlacklist struct {
	blacklist_repository repository.BlackListRepositoryInterface
}

func NewCheckBlacklist(
	blacklist_repository repository.BlackListRepositoryInterface) *UsecaseCheckBlacklist {
	return &UsecaseCheckBlacklist{blacklist_repository: blacklist_repository}
}

func (c *UsecaseCheckBlacklist) Execute(userIdentifier int, eventId string) (bool, string) {
	return c.blacklist_repository.Check(userIdentifier, eventId)
}
