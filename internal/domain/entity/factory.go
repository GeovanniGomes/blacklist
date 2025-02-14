package entity

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type FactoryEntity struct{}

func (f *FactoryEntity) FactoryNewBlacklist(eventId *string, reason, document, scope, blockedType string, userIdentifier int,isActive bool, blockedUntil, createdAt *time.Time, id *string)(*BlackList, error){
	valueId, err := f.setId(id)
	valueCreatedAt := f.setCreatedAt(createdAt)

	if err != nil{
		return  &BlackList{}, err
	}
	return NewBlackList(eventId,reason,document,scope,blockedType,userIdentifier,blockedUntil,valueCreatedAt, *valueId,isActive), nil

}

func (f *FactoryEntity) setId(id *string) (*string, error){
	newValue := uuid.NewV4().String()
	if id !=nil {
		value, err:= uuid.FromString(*id)
		if err !=nil{
			return nil, errors.New("value id invalid")
		}
		strValue := value.String()
		return &strValue, nil
	}
	return &newValue, nil
}

func (f *FactoryEntity) setCreatedAt(createdAt *time.Time) time.Time {
	newCreatedAt := time.Now()
	if createdAt !=nil {
		return *createdAt
	}
	return newCreatedAt
}