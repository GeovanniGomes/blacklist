package application

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestCheckBlackList(t *testing.T) {
	repositoryMemory := blacklist.BlackListRepositoryMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory)
	usecaseCheckBacklist := usecase.NewCheckBlacklist(&repositoryMemory)

	blacklistEntity, err := usecaseAddBacklist.Execute(10, uuid.NewV4().String(), "Fraude detectada", "10101010101", entity.GLOBAL, nil)
	require.Nil(t, err)

	message, err := usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), blacklistEntity.GetEventId())
	require.Nil(t, err)
	require.NotNil(t, message)
	require.Equal(t, message, "Fraude detectada")

	message, err = usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), uuid.NewV4().String())
	require.Nil(t, err)
	require.Equal(t, message, "")
}
