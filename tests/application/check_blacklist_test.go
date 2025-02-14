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
	blacklistEntity, err := usecaseAddBacklist.Execute(10, nil, "Fraude detectada", "10101010101", nil)
	require.Nil(t, err)
	require.Equal(t, blacklistEntity.GetScope(), entity.GLOBAL)
	message, err := usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), blacklistEntity.GetEventId())
	require.Nil(t, err)
	require.NotNil(t, message)
	require.Equal(t, message, "Fraude detectada")

	newEventId := uuid.NewV4().String()
	newBlacklist, _ := usecaseAddBacklist.Execute(blacklistEntity.GetUserIdentifier(), &newEventId, "Fraude detectada 2", "10101010101", nil)

	require.Equal(t, newBlacklist.GetScope(), entity.SPECIFIC)
	message, err = usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), &newEventId)
	require.Nil(t, err)
	require.Equal(t, message, "Fraude detectada 2")

	newEventIdTwo := uuid.NewV4().String()
	message, err = usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), &newEventIdTwo)
	require.Nil(t, err)
	require.Equal(t, message, "")
}
