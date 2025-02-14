package blacklist

import (
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

type IAddBlacklist interface {
	Execute(userIndentifier int, evendId *string, reason, document string, blockedUntil *time.Time) (*entity.BlackList, error)
}
