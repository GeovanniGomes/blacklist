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

func TestAddBlackList(t *testing.T){
	repositoryMemory := blacklist_repository.BlackListRepositoryMemory{}
	register_audit:= audit_repository.AuditLoggerMemory{}
	usecaseAddBacklist := usecase.NewAddBlacklist(&repositoryMemory, &register_audit)

	blacklistEntity , err := usecaseAddBacklist.Execute(10,uuid.NewV4().String(),"Fraude detectada","10101010101",entity.GLOBAL, nil)
	require.Nil(t,err)
	require.NotNil(t, blacklistEntity)
	require.Equal(t, blacklistEntity.GetScope(), entity.GLOBAL)
	require.Equal(t, blacklistEntity.GetBlockedType(), entity.PERMANENT)
}
