package application

import (
	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestReemoveBlackList(t *testing.T){
	repositoryMemory := persistence.BlackListRepositoryMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory)
	usecaseRemoveBacklist := usecase.NewRemoveBlacklist(&repositoryMemory)
	eventId := uuid.NewV4().String()
	eventIdTwo := uuid.NewV4().String()
	
	blacklistEntity , err := usecaseAddBacklist.Execute(10, eventId,"Fraude detectada","10101010101",entity.GLOBAL, nil)
	require.Nil(t,err)
	_ , err = usecaseAddBacklist.Execute(10,eventIdTwo,"Fraude no cartao","10101010101",entity.GLOBAL, nil)
	require.Nil(t,err)
	
	err = usecaseRemoveBacklist.Execute(blacklistEntity.GetUserIdentifier(), blacklistEntity.GetEventId())
	require.Nil(t, err)

}