package repository

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory/audit"
	"github.com/GeovanniGomes/blacklist/tests"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)


func TestAddLogger(t *testing.T){
	interface_database, teardown := tests.SetupPostgresContainer(t)
	defer teardown()

	repositoryBlacklist := audit.NewDBAuditLogger(interface_database)
	prepareBlacklist := entity.NewBlackList(uuid.NewV4().String(),"Fradude identificada","email@gmail.com", entity.GLOBAL, entity.PERMANENT,10, nil)

	err := prepareBlacklist.IsValid()
	require.Nil(t, err)
	logDetails := map[string]interface{}{
		"scope":         prepareBlacklist.GetScope(),
		"blocked_type":  prepareBlacklist.GetBlockedType(),
		"blocked_until": prepareBlacklist.GetBlockedUntil(),
	}
	repositoryBlacklist.LogAction(10,uuid.NewV4().String(),"add blacklist", logDetails)

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
