package entity

import (
	"github.com/GeovanniGomes/blacklist/internal/util"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	GLOBAL    = "global"
	SPECIFIC  = "specific"
	TEMPORARY = "temporary"
	PERMANENT = "permanent"
)

type BlackList struct {
	id             string
	eventId        string
	createdAt      time.Time
	reason         string
	document       string
	scope          string
	userIdentifier int
	blockedUntil   *time.Time
	blockedYype    string
}

func NewBlackList(enventId, reason, document, scope, blockedType string, userIdentifier int ,blockedUntil *time.Time) *BlackList {
	
	return &BlackList{
		id:             uuid.NewV4().String(),
		eventId:        enventId,
		reason:         reason,
		document:       document,
		scope:          scope,
		createdAt:      time.Now(),
		userIdentifier: userIdentifier,
		blockedUntil:   blockedUntil,
		blockedYype:    blockedType,
	}
}

func (blackList *BlackList) IsValid() error {
	now := util.TruncateDate(time.Now())
	if blackList.blockedUntil != nil {
		if util.TruncateDate(*blackList.blockedUntil).Before(now) {
			return errors.New("it is not possible to enable an event with a past date")
		}
	}

	if blackList.scope != GLOBAL && blackList.scope != SPECIFIC {
		return errors.New("scope must be global or specific")
	}
	if blackList.blockedYype != TEMPORARY && blackList.blockedYype != PERMANENT {
		return errors.New("blocked type must be temporary or permanent")
	}
	if blackList.blockedYype == TEMPORARY && blackList.blockedUntil == nil {
		return errors.New("blocked until is required for temporary block")
	}
	if util.GetSizeString(blackList.document) == 0 {
		return errors.New("document is required")
	}
	if util.GetSizeString(blackList.reason) == 0 {
		return errors.New("reason is required")
	}
	if blackList.userIdentifier == 0 {
		return errors.New("user identifier invalid")
	}

	_, err := uuid.FromString(blackList.eventId)
	if err != nil {
		return errors.New("event id is not a valid uuid")
	}

	return nil
}

func (blackList *BlackList) GetId() string {
	return blackList.id
}
func (blackList *BlackList) GetEventId() string {
	return blackList.eventId
}

func (blackList *BlackList) GetReason() string {
	return blackList.reason
}
func (blackList *BlackList) GetDocument() string {
	return blackList.document
}
func (blackList *BlackList) GetScope() string {
	return blackList.scope
}
func (blackList *BlackList) GetUserIdentifier() int {
	return blackList.userIdentifier
}
func (blackList *BlackList) GetBlockedUntil() *time.Time {
	return blackList.blockedUntil
}

func (blackList *BlackList) GetBlockedType() string {
	return blackList.blockedYype
}

func (blackList *BlackList) GetCreatedAt() time.Time {
	return blackList.createdAt
}
