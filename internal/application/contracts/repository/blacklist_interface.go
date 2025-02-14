package repositoty

import (
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

type IBlackListRepository interface {
	Check(userIndentifier int, eventId *string) (*entity.BlackList, error)
	Add(blacklist *entity.BlackList) error
	Remove(userIndentifier int, eventId string) error
	FetchBlacklistEntries(startDate, endDate time.Time) ([]entity.BlackList, error)
}
