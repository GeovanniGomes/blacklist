package entity_test

import (
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/tests/fixtures"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestBlackList_NewBlackList(t *testing.T) {
	dateEvent := time.Now()
	enventId := uuid.NewV4().String()
	newBlackList := fixtures.CreateBlacklist(10,enventId,"reason", "document",entity.GLOBAL, entity.TEMPORARY, &dateEvent)
	err:= newBlackList.IsValid()
	require.Nil(t, err)

	require.NotNil(t, newBlackList.GetId())
	require.Equal(t, newBlackList.GetEventId(),enventId)
	require.Equal(t, newBlackList.GetReason(), "reason")
	require.Equal(t, newBlackList.GetDocument(), "document")
	require.Equal(t, newBlackList.GetScope(), entity.GLOBAL)
	require.Equal(t, newBlackList.GetBlockedType(), entity.TEMPORARY)
	require.Equal(t, newBlackList.GetUserIdentifier(), 10)
	require.Equal(t, newBlackList.GetBlockedUntil(), &dateEvent)
	require.NotNil(t, newBlackList.GetCreatedAt())
}

func TestBlackList_NewBlackList_with_blockedUntil_invalid(t *testing.T) {
	dateEvent := time.Now()
	enventId := uuid.NewV4().String()
	dateEvent = dateEvent.AddDate(0,0,-1)
	newBlackList := fixtures.CreateBlacklist(10,enventId,"reason", "document",entity.GLOBAL, entity.TEMPORARY, &dateEvent)
	
	err:= newBlackList.IsValid()
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "it is not possible to enable an event with a past date")
}

func TestBlackList_NewBlackList_with_scope_invalid(t *testing.T) {
	dateEvent := time.Now()
	enventId := uuid.NewV4().String()
	newBlackList := fixtures.CreateBlacklist(10,enventId,"reason", "document","invalid", entity.TEMPORARY, &dateEvent)
	err:= newBlackList.IsValid()
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "scope must be global or specific")
}

func TestBlackList_NewBlackList_with_blockedType_invalid(t *testing.T) {
	dateEvent := time.Now()
	enventId := uuid.NewV4().String()
	newBlackList := fixtures.CreateBlacklist(10,enventId,"reason", "document",entity.GLOBAL, "invalid", &dateEvent)
	err := newBlackList.IsValid()
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "blocked type must be temporary or permanent")
}

func TestBlackList_NewBlackList_temporary_without_blockedutil(t *testing.T) {
	enventId := uuid.NewV4().String()
	newBlackList := fixtures.CreateBlacklist(10,enventId,"reason", "document",entity.GLOBAL, entity.TEMPORARY,nil)
	err := newBlackList.IsValid()
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "blocked until is required for temporary block")
}