package application

import (
	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)


func TestCheckBlackList(t *testing.T){
	repositoryMemory := persistence.BlackListRepositoryMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory)
	usecaseCheckBacklist := usecase.NewCheckBlacklist(&repositoryMemory)

	blacklistEntity , err := usecaseAddBacklist.Execute(10,uuid.NewV4().String(),"Fraude detectada","10101010101",entity.GLOBAL, nil)
	require.Nil(t,err)

	result, message := usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), blacklistEntity.GetEventId())
	require.NotNil(t,result)
	require.NotNil(t,message)
	require.Equal(t, result, false)
	require.Equal(t, message, "Fraude detectada")


	result, message = usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), uuid.NewV4().String())
	require.Equal(t, result, true)
	require.Equal(t, message, "")
}