package fixtures

import (
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/domain/value_objects"
	uuid "github.com/satori/go.uuid"
)

func CreateBlacklist(userIdentifier int, eventId *string, reason, document, scope, blockedType string, blockedUntil *time.Time) *entity.BlackList {
	return entity.NewBlackList(eventId, reason, document, scope, blockedType, userIdentifier, blockedUntil, time.Now(), uuid.NewV4().String(), true)
}

func CreateCategory(categoryName string) *value_objects.Category {
	category := value_objects.Category{}
	newCategory, _ := category.NewCategory(categoryName)
	return newCategory
}