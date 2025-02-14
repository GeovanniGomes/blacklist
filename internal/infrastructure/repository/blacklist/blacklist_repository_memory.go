package blacklist

import (
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

var _ repositoty.IBlackListRepository = (*BlackListRepositoryMemory)(nil)

type BlackListRepositoryMemory struct {
	collection_blacklist []entity.BlackList
}

func (black_list_repository *BlackListRepositoryMemory) FetchBlacklistEntries(startDate time.Time, endDate time.Time) ([]entity.BlackList, error) {
	var newCollection = []entity.BlackList{}
	for _, blacklist := range black_list_repository.collection_blacklist {
		if (blacklist.GetCreatedAt().After(startDate) || blacklist.GetCreatedAt().Equal(startDate)) && (blacklist.GetCreatedAt().Before(endDate) || blacklist.GetCreatedAt().Equal(endDate)) {
			newCollection = append(newCollection, blacklist)
		}
	}
	return newCollection, nil
}

func (black_list_repository *BlackListRepositoryMemory) Add(blacklist *entity.BlackList) error {
	black_list_repository.collection_blacklist = append(black_list_repository.collection_blacklist, *blacklist)
	return nil
}

func (black_list_repository *BlackListRepositoryMemory) Check(userIndentifier int, evendId *string) (*entity.BlackList, error) {
	for _, blacklist := range black_list_repository.collection_blacklist {
		if blacklist.GetUserIdentifier() == userIndentifier && blacklist.GetEventId() == evendId {
			return &blacklist, nil
		}
	}
	return &entity.BlackList{}, nil
}

func (black_list_repository *BlackListRepositoryMemory) Remove(userIndentifier int, eventId string) error {
	var newCollection = []entity.BlackList{}
	for _, blacklist := range black_list_repository.collection_blacklist {
		if !(blacklist.GetUserIdentifier() == userIndentifier && blacklist.GetEventId() == &eventId) {
			newCollection = append(newCollection, blacklist)
			continue
		}
	}
	black_list_repository.collection_blacklist = newCollection
	return nil
}
