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

func TestReemoveBlackList(t *testing.T) {
	repositoryMemory := blacklist_repository.BlackListRepositoryMemory{}
	auditRepositoryMemory := audit_repository.AuditLoggerMemory{}

	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory, &auditRepositoryMemory)
	usecaseRemoveBacklist := usecase.NewRemoveBlacklist(&repositoryMemory)
	eventId := uuid.NewV4().String()
	eventIdTwo := uuid.NewV4().String()

	blacklistEntity, err := usecaseAddBacklist.Execute(10, eventId, "Fraude detectada", "10101010101", entity.GLOBAL, nil)
	require.Nil(t, err)
	_, err = usecaseAddBacklist.Execute(10, eventIdTwo, "Fraude no cartao", "10101010101", entity.GLOBAL, nil)
	require.Nil(t, err)

	err = usecaseRemoveBacklist.Execute(blacklistEntity.GetUserIdentifier(), blacklistEntity.GetEventId())
	require.Nil(t, err)

}
