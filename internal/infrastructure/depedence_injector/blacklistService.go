package depedence_injector

import (
	usecaseBlacklistContract "github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	contracts_infrastructure "github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/producer"
	"go.uber.org/dig"
)

func RegistreBlackList(c *dig.Container){
	c.Provide(func(
		usecaseCreateBlacklist usecaseBlacklistContract.IAddBlacklist,
		usecaseCheckBlacklist usecaseBlacklistContract.ICheckBlacklist,
		usecaseRemoveBlacklist usecaseBlacklistContract.IRemoveBlackList,
		register_audit contracts_infrastructure.IAuditLogger,
		persistence_cache contracts_infrastructure.ICache,
		producer *producer.BlacklistProducer,
	) *service.BlacklistService {
		return service.NewBlackListService(
			usecaseCreateBlacklist,
			usecaseCheckBlacklist,
			usecaseRemoveBlacklist,
			register_audit,
			persistence_cache,
			producer,
		)
	})
}