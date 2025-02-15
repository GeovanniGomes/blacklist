package repository

import (
	"testing"
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/domain/value_objects"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"
	"github.com/GeovanniGomes/blacklist/tests"
	"github.com/stretchr/testify/require"
)

func TestSaveEvent(t *testing.T) {
	interface_database, teardown := tests.SetupPostgresContainer(t)
	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	defer teardown()

	factory := entity.FactoryEntity{}
	dateEvent := time.Now().Add(48)
	prepareEvent, err := factory.FactoryNewEvent(nil, "Jogo do flamento", "Brasileirão", dateEvent, nil, value_objects.SOCCER, true)
	require.Nil(t, err)

	err = prepareEvent.IsValid()
	require.Nil(t, err)
	err = repositoryBlacklist.AddEvent(*prepareEvent)
	require.Nil(t, err)

	event, err := repositoryBlacklist.GetEvent(prepareEvent.GetId())
	require.Nil(t, err)
	require.Equal(t, event.GetId(), prepareEvent.GetId())

}
