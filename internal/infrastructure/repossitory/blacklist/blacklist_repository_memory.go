package blacklist

import (
	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

var _ repositoty.IBlackListRepository = (*BlackListRepositoryMemory)(nil)

type BlackListRepositoryMemory struct {
	collection_blacklist []entity.BlackList
}

func (black_list_repository *BlackListRepositoryMemory) Add(blacklist *entity.BlackList) error {
	black_list_repository.collection_blacklist = append(black_list_repository.collection_blacklist, *blacklist)
	return nil
}

func (black_list_repository *BlackListRepositoryMemory) Check(userIndentifier int, evendId string) (bool, string) {
	for _, blacklist := range black_list_repository.collection_blacklist {
		if blacklist.GetUserIdentifier() == userIndentifier && blacklist.GetEventId() == evendId {
			return false, blacklist.GetReason()
		}
	}
	return true, ""
}

func (black_list_repository *BlackListRepositoryMemory) Remove(userIndentifier int, eventId string) error {
	var newCollection = []entity.BlackList{}
	for _, blacklist := range black_list_repository.collection_blacklist {
		if !(blacklist.GetUserIdentifier() == userIndentifier && blacklist.GetEventId() == eventId) {
			newCollection = append(newCollection, blacklist)
			continue
		}
	}
	black_list_repository.collection_blacklist = newCollection
	return nil
}
