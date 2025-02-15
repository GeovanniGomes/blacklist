package blacklist

import (
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

var _ repositoty.IBlackListRepository = (*BlackListRepositoryMemory)(nil)

type BlackListRepositoryMemory struct {
	collectionBlacklist []entity.BlackList
	collectionEvents    []entity.Event
}

func (black_list_repository *BlackListRepositoryMemory) FetchBlacklistEntries(startDate time.Time, endDate time.Time) ([]entity.BlackList, error) {
	var newCollection = []entity.BlackList{}
	for _, blacklist := range black_list_repository.collectionBlacklist {
		if (blacklist.GetCreatedAt().After(startDate) || blacklist.GetCreatedAt().Equal(startDate)) && (blacklist.GetCreatedAt().Before(endDate) || blacklist.GetCreatedAt().Equal(endDate)) {
			newCollection = append(newCollection, blacklist)
		}
	}
	return newCollection, nil
}

func (black_list_repository *BlackListRepositoryMemory) AddBlacklist(blacklist *entity.BlackList) error {
	black_list_repository.collectionBlacklist = append(black_list_repository.collectionBlacklist, *blacklist)
	return nil
}

func (black_list_repository *BlackListRepositoryMemory) CheckBlacklist(userIndentifier int, evendId *string) (*entity.BlackList, error) {
	for _, blacklist := range black_list_repository.collectionBlacklist {
		if blacklist.GetUserIdentifier() == userIndentifier && blacklist.GetEventId() == evendId {
			return &blacklist, nil
		}
	}
	return &entity.BlackList{}, nil
}

func (black_list_repository *BlackListRepositoryMemory) RemoveBlacklist(userIndentifier int, eventId string) error {
	var newCollection = []entity.BlackList{}
	for _, blacklist := range black_list_repository.collectionBlacklist {
		if !(blacklist.GetUserIdentifier() == userIndentifier && blacklist.GetEventId() == &eventId) {
			newCollection = append(newCollection, blacklist)
			continue
		}
	}
	black_list_repository.collectionBlacklist = newCollection
	return nil
}

func (black_list_repository *BlackListRepositoryMemory) AddEvent(event entity.Event) error {
	black_list_repository.collectionEvents = append(black_list_repository.collectionEvents, event)
	return nil
}

func (black_list_repository *BlackListRepositoryMemory) GetEvent(id string) (*entity.Event, error) {
	for _, event := range black_list_repository.collectionEvents {
		if event.GetId() == id {
			return &event, nil
		}
	}
	return &entity.Event{}, nil
}

func (black_list_repository *BlackListRepositoryMemory) RemoveEvent(id string) error {
	var newCollection = []entity.Event{}
	for _, event := range black_list_repository.collectionEvents {
		if !(event.GetId() == id) {
			newCollection = append(newCollection, event)
			continue
		}
	}
	black_list_repository.collectionEvents = newCollection
	return nil
}
