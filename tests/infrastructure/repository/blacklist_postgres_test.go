package repository

import (
	"testing"
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"
	"github.com/GeovanniGomes/blacklist/tests"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestSaveBlackList(t *testing.T) {
	interface_database, teardown := tests.SetupPostgresContainer(t)
	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	defer teardown()

	factory := entity.FactoryEntity{}
	eventId := uuid.NewV4().String()
	prepareBlacklist, err := factory.FactoryNewBlacklist(&eventId, "Fradude identificada", "email@gmail.com", entity.SPECIFIC, entity.PERMANENT, 10, true, nil, nil, nil)
	require.Nil(t, err)

	err = prepareBlacklist.IsValid()
	require.Nil(t, err)
	repositoryBlacklist.Add(prepareBlacklist)

	rows, err := interface_database.SelectQuery("SELECT reason FROM blacklist WHERE id = $1", prepareBlacklist.GetId())
	require.Nil(t, err)
	defer rows.Close()

	var reason string
	if rows.Next() {
		err = rows.Scan(&reason)
		require.Nil(t, err)
		require.Equal(t, "Fradude identificada", reason)
	}
}

func TestSaveBlackListWithFetch(t *testing.T) {
	interface_database, teardown := tests.SetupPostgresContainer(t)
	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	defer teardown()

	factory := entity.FactoryEntity{}
	startDate := time.Now().Truncate(24 * time.Hour)
	endDate := startDate.Add(24 * time.Hour)

	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, endDate.Location())

	eventId := uuid.NewV4().String()
	prepareBlacklist, err := factory.FactoryNewBlacklist(&eventId, "Fradude identificada", "email@gmail.com", entity.SPECIFIC, entity.PERMANENT, 10, true, nil, nil, nil)
	require.Nil(t, err)

	err = prepareBlacklist.IsValid()
	require.Nil(t, err)
	err = repositoryBlacklist.Add(prepareBlacklist)

	require.Nil(t, err)

	rowsBlacklist, err := repositoryBlacklist.FetchBlacklistEntries(startDate, endDate)
	require.Nil(t, err)
	require.NotNil(t, rowsBlacklist)

}

func TestCheckBlackList(t *testing.T) {
	interface_database, teardown := tests.SetupPostgresContainer(t)
	defer teardown()
	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)

	factory := entity.FactoryEntity{}
	eventId := uuid.NewV4().String()
	prepareBlacklist, err := factory.FactoryNewBlacklist(&eventId, "Fradude identificada", "email@gmail.com", entity.SPECIFIC, entity.PERMANENT, 10, true, nil, nil, nil)
	require.Nil(t, err)

	err = prepareBlacklist.IsValid()
	require.Nil(t, err)
	repositoryBlacklist.Add(prepareBlacklist)
	blacklist, err := repositoryBlacklist.Check(prepareBlacklist.GetUserIdentifier(), prepareBlacklist.GetEventId())

	require.Nil(t, err)
	require.Equal(t, blacklist.GetReason(), "Fradude identificada")
}

func TestRemoveBlackList(t *testing.T) {
	interface_database, teardown := tests.SetupPostgresContainer(t)
	factory := entity.FactoryEntity{}

	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	defer teardown()
	eventId := uuid.NewV4().String()
	prepareBlacklist, err := factory.FactoryNewBlacklist(&eventId, "Fradude identificada", "email@gmail.com", entity.SPECIFIC, entity.PERMANENT, 10, true, nil, nil, nil)
	require.Nil(t, err)
	err = prepareBlacklist.IsValid()
	require.Nil(t, err)
	repositoryBlacklist.Add(prepareBlacklist)

	err = repositoryBlacklist.Remove(prepareBlacklist.GetUserIdentifier(), *prepareBlacklist.GetEventId())
	require.Nil(t, err)
	blacklist, err := repositoryBlacklist.Check(prepareBlacklist.GetUserIdentifier(), prepareBlacklist.GetEventId())

	require.Nil(t, err)
	require.Equal(t, blacklist.GetReason(), "")
}
