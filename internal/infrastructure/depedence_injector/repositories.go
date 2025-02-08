package depedence_injector

import (
	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/audit"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"

	"go.uber.org/dig"
)

func RegisterRepository(c *dig.Container) {
	c.Provide(func(persistence contracts.IDatabaseRelational) repositoty.IBlackListRepository {
		return blacklist.NewBlackListRepositoryPostgres(persistence)
	})

	c.Provide(func(persistence contracts.IDatabaseRelational) contracts.IAuditLogger {
		return audit.NewDBAuditLogger(persistence)
	})
	
	c.Provide(func(persistence contracts.IDatabaseRelational) repositoty.IBlackListRepository {
		return blacklist.NewBlackListRepositoryPostgres(persistence)
	})
}