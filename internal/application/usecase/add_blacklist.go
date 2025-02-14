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

func (c *UsecaseAddBlacklist) Execute(userIdentifier int, eventId *string, reason, document string, blocked_until *time.Time) (*entity.BlackList, error) {
	blacklistEmpty := entity.BlackList{}
	scope := entity.GLOBAL
	factory := entity.FactoryEntity{}
	blocked_type := entity.TEMPORARY
	if blocked_until == nil {
		blocked_type = entity.PERMANENT
	}

	if eventId != nil {
		scope = entity.SPECIFIC
	}

	blacklist, err := factory.FactoryNewBlacklist(eventId, reason,document,scope,blocked_type,userIdentifier,true,blocked_until,nil,nil)
	if err !=nil{
		log.Printf("error factory entity blacklist %v", err)
		return &blacklistEmpty, errors.New("unable to add blacklist")
	}
	if err = blacklist.IsValid(); err != nil {
		return &blacklistEmpty, err
	}
	
	if err = c.blacklist_repository.AddBlacklist(blacklist); err != nil {
		return &blacklistEmpty, err
	}

	return blacklist, nil
}
