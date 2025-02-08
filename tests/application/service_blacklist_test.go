package application

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/application/dto"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/producer"
	repository_redis "github.com/GeovanniGomes/blacklist/internal/infrastructure/repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/audit"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"
	"github.com/GeovanniGomes/blacklist/tests"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func SetupRunTest(t *testing.T) (*service.BlacklistService, contracts.IDatabaseRelational, contracts.ICache) {
	interface_database, teardown := tests.SetupPostgresContainer(t)
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

	url := os.Getenv("CONNECTION_STRING_BROKEN_QUEUE")
	rabbitQueue := queue.NewRabbitMQQueue(url)
	dispatcher := queue.NewDispatcher(rabbitQueue)
	blacklistProducer := producer.NewBlacklistProducer(dispatcher)

	return service.NewBlackListService(
		usecaseAddBlacklist,
		usecaseCheckBlacklist,
		usecaseRemoveBlacklist,
		register_audit, persistence_cache,
		blacklistProducer,
	), interface_database, persistence_cache

}
func TestAddBlackListService(t *testing.T) {
	blacklitervice, interface_database, _ := SetupRunTest(t)
	eventId := uuid.NewV4().String()
	requestInput := dto.BlacklistInput{
		EventId:        eventId,
		UserIdentifier: 10,
		Scope:          "global",
		BlockedUntil:   nil,
		Reason:         "Nao paga mensalidade",
		Document:       "101101101101",
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
	eventId := uuid.NewV4().String()
	userIdentifier := 10
	requestInput := dto.BlacklistInput{
		EventId:        eventId,
		UserIdentifier: userIdentifier,
		Scope:          "global",
		BlockedUntil:   nil,
		Reason:         "Nao paga mensalidade",
		Document:       "101101101101",
	}
	err := blacklitervice.AddBlacklist(requestInput)
	require.Nil(t, err)

	requestInputCheck := dto.BlacklistInputCheck{
		UserIdentifier: userIdentifier,
		EventId:        eventId,
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
	eventId := uuid.NewV4().String()
	userIdentifier := 10
	requestInput := dto.BlacklistInput{
		EventId:        eventId,
		UserIdentifier: userIdentifier,
		Scope:          "global",
		BlockedUntil:   nil,
		Reason:         "Nao paga mensalidade",
		Document:       "101101101101",
	}
	err := blacklitervice.AddBlacklist(requestInput)
	require.Nil(t, err)

	requestInputRemove := dto.BlacklistInputRemove{
		UserIdentifier: userIdentifier,
		EventId:        eventId,
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
	require.Equal(t, fmt.Sprintf("key %v not found in cache", key), err.Error())
}
