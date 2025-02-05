package application

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory/blacklist"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestAddBlackList(t *testing.T){
	repositoryMemory := blacklist.BlackListRepositoryMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory)

	blacklistEntity , err := usecaseAddBacklist.Execute(10,uuid.NewV4().String(),"Fraude detectada","10101010101",entity.GLOBAL, nil)
	require.Nil(t,err)
	require.NotNil(t, blacklistEntity)
	require.Equal(t, blacklistEntity.GetScope(), entity.GLOBAL)
	require.Equal(t, blacklistEntity.GetBlockedType(), entity.PERMANENT)
}
