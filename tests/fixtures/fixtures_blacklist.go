package fixtures

import (
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"time"
)

func CreateBlacklist(userIdentifier int, eventId, reason, document, scope, blockedType string, blockedUntil *time.Time )(*entity.BlackList){
	return entity.NewBlackList(eventId,reason,document,scope,blockedType,userIdentifier,blockedUntil)
}