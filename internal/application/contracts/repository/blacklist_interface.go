package repositoty

import (
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

type IBlackListRepository interface {
	CheckBlacklist(userIndentifier int, eventId *string) (*entity.BlackList, error)
	AddBlacklist(blacklist *entity.BlackList) error
	RemoveBlacklist(userIndentifier int, eventId string) error
	FetchBlacklistEntries(startDate, endDate time.Time) ([]entity.BlackList, error)

	AddEvent(event entity.Event) error
	GetEvent(id string) (*entity.Event,error)
}
