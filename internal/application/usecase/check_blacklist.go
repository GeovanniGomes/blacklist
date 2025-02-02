package usecase

import (
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

var _ blacklist.CheckBlacklistInterface = (*UsecaseCheckBlacklist)(nil)

type UsecaseCheckBlacklist struct {
	blacklist_repository repository.BlackListRepositoryInterface
	register_audit       contracts.AuditLoggerInterface
}

func NewCheckBlacklist(
	blacklist_repository repository.BlackListRepositoryInterface, register_audit contracts.AuditLoggerInterface) *UsecaseCheckBlacklist {
	return &UsecaseCheckBlacklist{blacklist_repository: blacklist_repository, register_audit: register_audit}
}

func (c *UsecaseCheckBlacklist) Execute(userIdentifier int, eventId string) (bool, string) {
	result, reason := c.blacklist_repository.Check(userIdentifier, eventId)
	logDetails := map[string]interface{}{
		"is_blocked": result,
		"reason":     reason,
	}
	c.register_audit.LogAction(userIdentifier, eventId, contracts.CHECK_BLACKLIST, &logDetails)
	return result, reason
}
