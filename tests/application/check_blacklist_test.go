package application

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence/audit_repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence/blacklist_repository"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestCheckBlackList(t *testing.T) {
	repositoryMemory := blacklist_repository.BlackListRepositoryMemory{}
	auditRepositoryMemory := audit_repository.AuditLoggerMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory, &auditRepositoryMemory)
	usecaseCheckBacklist := usecase.NewCheckBlacklist(&repositoryMemory, &auditRepositoryMemory)

	blacklistEntity, err := usecaseAddBacklist.Execute(10, uuid.NewV4().String(), "Fraude detectada", "10101010101", entity.GLOBAL, nil)
	require.Nil(t, err)

	result, message := usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), blacklistEntity.GetEventId())
	require.NotNil(t, result)
	require.NotNil(t, message)
	require.Equal(t, result, false)
	require.Equal(t, message, "Fraude detectada")

	result, message = usecaseCheckBacklist.Execute(blacklistEntity.GetUserIdentifier(), uuid.NewV4().String())
	require.Equal(t, result, true)
	require.Equal(t, message, "")
}
