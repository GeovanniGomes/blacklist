package usecase

import (
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

var _ blacklist.AddBlacklistInterface = (*UsecaseAddBlacklist)(nil)

type UsecaseAddBlacklist struct {
	blacklist_repository repositoty.BlackListRepositoryInterface
	register_audit       contracts.AuditLoggerInterface
}

func NewAddBlacklist(blacklist_repository repositoty.BlackListRepositoryInterface, register_audit contracts.AuditLoggerInterface) *UsecaseAddBlacklist {
	return &UsecaseAddBlacklist{blacklist_repository: blacklist_repository, register_audit: register_audit}
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
	logDetails := map[string]interface{}{
		"scope":         scope,
		"blocked_type":  blocked_type,
		"blocked_until": blocked_until,
	}
	c.register_audit.LogAction(userIdentifier, eventId, contracts.ADD_BLACKLIST, &logDetails)
	return blacklist, nil
}
