package entity

import (
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/value_objects"
	"github.com/GeovanniGomes/blacklist/internal/util"
)

type FactoryEntity struct{}

func (f *FactoryEntity) FactoryNewBlacklist(eventId *string, reason, document, scope, blockedType string, userIdentifier int,isActive bool, blockedUntil, createdAt *time.Time, id *string)(*BlackList, error){
	valueId, err := util.ValidateOrGenerateID(id)
	valueCreatedAt := util.DefaultOrProvidedTime(createdAt)

	if err != nil{
		return  &BlackList{}, err
	}
	return NewBlackList(eventId,reason,document,scope,blockedType,userIdentifier,blockedUntil,valueCreatedAt, *valueId,isActive), nil

}
func (f *FactoryEntity) FactoryNewEvent(id *string,title, description string, date time.Time, createdAt *time.Time, category string, isActive bool, status string)(*Event, error){
	valueId, err :=  util.ValidateOrGenerateID(id)
	valueCreatedAt :=util.DefaultOrProvidedTime(createdAt)

	if err != nil{
		return  &Event{}, err
	}
	categoryValueObject := value_objects.Category{}
	categoryEntity, err := categoryValueObject.NewCategory(category)
	if err != nil{
		return &Event{},err
	}
	return NewEvent(*valueId,title,description,date,valueCreatedAt,*categoryEntity,isActive, status), nil

}