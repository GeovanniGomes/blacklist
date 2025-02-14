package application

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestAddBlackList(t *testing.T){
	repositoryMemory := blacklist.BlackListRepositoryMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory)
	eventId:= uuid.NewV4().String()
	blacklistEntity , err := usecaseAddBacklist.Execute(10,&eventId,"Fraude detectada","10101010101", nil)
	require.Nil(t,err)
	require.NotNil(t, blacklistEntity)
	require.Equal(t, blacklistEntity.GetScope(), entity.SPECIFIC)
	require.Equal(t, blacklistEntity.GetBlockedType(), entity.PERMANENT)
}
