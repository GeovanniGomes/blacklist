package container_repository

import (
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/audit"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"
	"go.uber.org/dig"
)

func RegisterRepository(c *dig.Container) {
	c.Provide(func(persistence contracts.IDatabaseRelational) repository.IBlackListRepository {
		return blacklist.NewBlackListRepositoryPostgres(persistence)
	})

	c.Provide(func(persistence contracts.IDatabaseRelational) contracts.IAuditLogger {
		return audit.NewDBAuditLogger(persistence)
	})

	c.Provide(func(persistence contracts.IDatabaseRelational) repository.IBlackListRepository {
		return blacklist.NewBlackListRepositoryPostgres(persistence)
	})
}
