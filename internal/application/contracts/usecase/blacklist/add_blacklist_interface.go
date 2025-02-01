package blacklist

import (
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"time"
)

type AddBlacklistInterface interface {
	Execute(userIndentifier int, evendId, reason, document, scope string, blockedUntil *time.Time ) (*entity.BlackList, error)
}