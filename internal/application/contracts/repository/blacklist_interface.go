package repositoty

import "github.com/GeovanniGomes/blacklist/internal/domain/entity"


type BlackListRepositoryInterface interface {
	Check(userIndentifier int, evendId string) (bool, string)
	Add(blacklist *entity.BlackList) error
	Remove(userIndentifier int, eventId string) error
}
