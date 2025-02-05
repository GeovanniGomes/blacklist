package application

import (
	"context"
	"fmt"
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/application/dto"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	repository_redis "github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory/audit"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory/blacklist"
	"github.com/GeovanniGomes/blacklist/tests/infrastructure"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func SetupRunTest(t *testing.T) (*service.BlacklistService, contracts.DatabaseRelationalInterface, contracts.CacheInterface) {
	interface_database, teardown := infrastructure.SetupPostgresContainer(t)
	defer teardown()

	repositoryBlacklist := blacklist.NewBlackListRepositoryPostgres(interface_database)
	usecaseAddBlacklist := usecase.NewAddBlacklist(repositoryBlacklist)
	usecaseCheckBlacklist := usecase.NewCheckBlacklist(repositoryBlacklist)
	usecaseRemoveBlacklist := usecase.NewRemoveBlacklist(repositoryBlacklist)

	register_audit := audit.NewDBAuditLogger(interface_database)
	persistence_cache, err := repository_redis.NewRedisService("localhost:6379", "", 1)
	if err != nil {
		t.Fatal(err)
	}

	return service.NewBlackListService(
		usecaseAddBlacklist,
		usecaseCheckBlacklist,
		usecaseRemoveBlacklist,
		register_audit, persistence_cache,
	), interface_database, persistence_cache
}
func TestAddBlackListService(t *testing.T) {
	blacklitervice, interface_database, _ := SetupRunTest(t)
	eventId :=  uuid.NewV4().String()
	requestInput := dto.BlacklistInput{
		EventId:        eventId,
		UserIdentifier: 10,
		Scope:          "global",
		BlockedUntil:   nil,
		Reason:         "Nao paga mensalidade",
		Document: "101101101101",
	}
	err := blacklitervice.AddBlacklist(requestInput)
	require.Nil(t, err)

	rows, err := interface_database.SelectQuery("SELECT reason FROM blacklist WHERE event_id = $1", eventId)
	require.Nil(t, err)
	defer rows.Close()

	var reason string
	if rows.Next() {
		err = rows.Scan(&reason)
		require.Nil(t, err)
		require.Equal(t, "Nao paga mensalidade", reason)
	}

}

func TestCheckBlackListService(t *testing.T) {
	ctx := context.Background()
	blacklitervice, _, persistence_cache := SetupRunTest(t)
	eventId :=  uuid.NewV4().String()
	userIdentifier :=10
	requestInput := dto.BlacklistInput{
		EventId:        eventId,
		UserIdentifier: userIdentifier,
		Scope:          "global",
		BlockedUntil:   nil,
		Reason:         "Nao paga mensalidade",
		Document: "101101101101",
	}
	err := blacklitervice.AddBlacklist(requestInput)
	require.Nil(t, err)

	requestInputCheck :=dto.BlacklistInputCheck{
		UserIdentifier: userIdentifier,
		EventId: eventId,
	} 
	result, err := blacklitervice.CheckBlacklist(requestInputCheck)
	require.Nil(t, err)

	key := fmt.Sprintf("%v_%v", userIdentifier, eventId)
	detailCache, err := persistence_cache.GetCache(ctx, key)
	require.Nil(t, err)
	

	require.Equal(t, result.IsBlocked, detailCache["is_blocked"])
	require.Equal(t, result.Reason, detailCache["reason"])
}

func TestRemokBlackListService(t *testing.T) {
	ctx := context.Background()
	blacklitervice, interface_database, persistence_cache := SetupRunTest(t)
	eventId :=  uuid.NewV4().String()
	userIdentifier :=10
	requestInput := dto.BlacklistInput{
		EventId:        eventId,
		UserIdentifier: userIdentifier,
		Scope:          "global",
		BlockedUntil:   nil,
		Reason:         "Nao paga mensalidade",
		Document: "101101101101",
	}
	err := blacklitervice.AddBlacklist(requestInput)
	require.Nil(t, err)

	requestInputRemove :=dto.BlacklistInputRemove{
		UserIdentifier: userIdentifier,
		EventId: eventId,
	} 
	err = blacklitervice.RemoveBlacklist(requestInputRemove)
	require.Nil(t, err)
	
	rows, err := interface_database.SelectQuery("SELECT reason FROM blacklist WHERE event_id = $1 and is_active = $2", eventId, true)
	require.Nil(t, err)
	defer rows.Close()

	require.Equal(t, rows.Next(), false)

	key := fmt.Sprintf("%v_%v", userIdentifier, eventId)
	_, err = persistence_cache.GetCache(ctx, key)
	require.NotNil(t, err)
	require.Equal(t, fmt.Sprintf("key %v not found in cache", key),err.Error())
}