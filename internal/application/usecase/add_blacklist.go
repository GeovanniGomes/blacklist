package usecase

import (
	"errors"
	"log"
	"time"

	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
)

var _ blacklist.IAddBlacklist = (*UsecaseAddBlacklist)(nil)

type UsecaseAddBlacklist struct {
	blacklist_repository repositoty.IBlackListRepository
}

func NewAddBlacklist(blacklist_repository repositoty.IBlackListRepository) *UsecaseAddBlacklist {
	return &UsecaseAddBlacklist{blacklist_repository: blacklist_repository}
}

func (c *UsecaseAddBlacklist) Execute(userIdentifier int, eventId, reason, document, scope string, blocked_until *time.Time) (*entity.BlackList, error) {
	var blacklistEmpty = entity.BlackList{}
	factory := entity.FactoryEntity{}
	blocked_type := entity.TEMPORARY
	if blocked_until == nil {
		blocked_type = entity.PERMANENT
	}
	blacklist, err := factory.FactoryNewBlacklist(eventId, reason,document,scope,blocked_type,userIdentifier,true,blocked_until,nil,nil)
	if err !=nil{
		log.Printf("error factory entity blacklist %v", err)
		return &blacklistEmpty, errors.New("unable to add blacklist")
	}
	
	err = blacklist.IsValid()
	
	if err != nil {
		return &blacklistEmpty, err
	}
	err = c.blacklist_repository.Add(blacklist)
	if err != nil {
		return &blacklistEmpty, err
	}
	return blacklist, nil
}
