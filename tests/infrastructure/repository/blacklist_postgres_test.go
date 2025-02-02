package repository

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory/blacklist"
	"github.com/GeovanniGomes/blacklist/tests/infrastructure"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)


func TestSaveBlackList(t *testing.T){
	interface_database, teardown := infrastructure.SetupPostgresContainer(t)
	defer teardown()

	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	prepareBlacklist := entity.NewBlackList(uuid.NewV4().String(),"Fradude identificada","email@gmail.com", entity.GLOBAL, entity.PERMANENT,10, nil)

	err := prepareBlacklist.IsValid()
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

func TestCheckBlackList(t *testing.T){
	interface_database, teardown := infrastructure.SetupPostgresContainer(t)
	defer teardown()

	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	prepareBlacklist := entity.NewBlackList(uuid.NewV4().String(),"Fradude identificada","email@gmail.com", entity.GLOBAL, entity.PERMANENT,10, nil)

	err := prepareBlacklist.IsValid()
	require.Nil(t, err)
	repositoryBlacklist.Add(prepareBlacklist)

	result, reason := repositoryBlacklist.Check(prepareBlacklist.GetUserIdentifier(), prepareBlacklist.GetEventId())

	require.Equal(t, result, false)
	require.Equal(t, reason, "Fradude identificada")
}

func TestRemoveBlackList(t *testing.T){
	interface_database, teardown := infrastructure.SetupPostgresContainer(t)
	defer teardown()

	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	prepareBlacklist := entity.NewBlackList(uuid.NewV4().String(),"Fradude identificada","email@gmail.com", entity.GLOBAL, entity.PERMANENT,10, nil)

	err := prepareBlacklist.IsValid()
	require.Nil(t, err)
	repositoryBlacklist.Add(prepareBlacklist)

	err = repositoryBlacklist.Remove(prepareBlacklist.GetUserIdentifier(), prepareBlacklist.GetEventId())
	require.Nil(t, err)
	result, reason := repositoryBlacklist.Check(prepareBlacklist.GetUserIdentifier(), prepareBlacklist.GetEventId())
	
	require.Equal(t, result, true)
	require.Equal(t, reason, "")
}