package fixtures

import (
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	uuid "github.com/satori/go.uuid"
)

func CreateBlacklist(userIdentifier int, eventId, reason, document, scope, blockedType string, blockedUntil *time.Time) *entity.BlackList {
	return entity.NewBlackList(eventId, reason, document, scope, blockedType, userIdentifier, blockedUntil, time.Now(), uuid.NewV4().String(), true)
}
