package depedence_injector

import (
	//repositoryBlacklistContract "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	//contracts_infrastructure "github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	//"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"

	//repository_providers_implementation "github.com/GeovanniGomes/blacklist/internal/infrastructure/repository"
	//repository_audit "github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/audit"
	//repository_blacklist "github.com/GeovanniGomes/blacklist/internal/infrastructure/repository/blacklist"

	usecaseBlacklistContract "github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	//usecseApplicationBlacklist "github.com/GeovanniGomes/blacklist/internal/application/usecase"

	"go.uber.org/dig"
)

type ContainerInjection struct {
	*dig.Container
}

func NewContainer() *ContainerInjection {
	c := dig.New()

	RegisterDatabase(c)
	RegisterQueue(c)
	RegisterDispatcher(c)
	RegisterCache(c)
	RegisterRepository(c)
	RegisterUsecase(c)
	RegistreBlackList(c)

	// // Register instancia database witch postrgres
	// c.Provide(func() contracts_infrastructure.IDatabaseRelational {
	// 	instance, err := repository_providers_implementation.NewPostgresDatabase(os.Getenv("CONNECTION_STRING_DATABASE"))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return instance
	// })

	// // regisster broken rabbitMQ  queue
	// c.Provide(func() contracts_infrastructure.IQueue {
	// 	return queue.NewRabbitMQQueue(os.Getenv("CONNECTION_STRING_BROKEN_QUEUE"))
	// })

	// c.Provide(func(handler contracts_infrastructure.IQueue) *queue.Dispatcher {
	// 	return queue.NewDispatcher(handler)
	// })

	// // register cache
	// c.Provide(func() contracts_infrastructure.ICache {
	// 	addr := os.Getenv("REDIS_ADDR")
	// 	password := os.Getenv("REDIS_PASSWORD")
	// 	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	// 	instance, err := repository_providers_implementation.NewRedisService(addr, password, db)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return instance
	// })

	// c.Provide(func(dispatcher *queue.Dispatcher) *producer.BlacklistProducer {
	// 	return producer.NewBlacklistProducer(dispatcher)
	// })

	// // Register repositories
	// // Register repository blacklist
	// c.Provide(func(persistence contracts_infrastructure.IDatabaseRelational) repositoryBlacklistContract.IBlackListRepository {
	// 	return repository_blacklist.NewBlackListRepositoryPostgres(persistence)
	// })

	// // Registra o repository audit
	// c.Provide(func(persistence contracts_infrastructure.IDatabaseRelational) contracts_infrastructure.IAuditLogger {
	// 	return repository_audit.NewDBAuditLogger(persistence)
	// })

	// // register repository blacklist
	// c.Provide(func(persistence contracts_infrastructure.IDatabaseRelational) repositoryBlacklistContract.IBlackListRepository {
	// 	return repository_blacklist.NewBlackListRepositoryPostgres(persistence)
	// })

	// // Register usecases add blackisst
	// c.Provide(func(repository repositoryBlacklistContract.IBlackListRepository) usecaseBlacklistContract.IAddBlacklist {
	// 	return usecseApplicationBlacklist.NewAddBlacklist(repository)
	// })

	// // Register usecases check blacklisst
	// c.Provide(func(repository repositoryBlacklistContract.IBlackListRepository) usecaseBlacklistContract.ICheckBlacklist {
	// 	return usecseApplicationBlacklist.NewCheckBlacklist(repository)
	// })

	// // Register usecases remove blacklisst
	// c.Provide(func(repository repositoryBlacklistContract.IBlackListRepository) usecaseBlacklistContract.IRemoveBlackList {
	// 	return usecseApplicationBlacklist.NewRemoveBlacklist(repository)
	// })

	// c.Provide(func(
	// 	usecaseCreateBlacklist usecaseBlacklistContract.IAddBlacklist,
	// 	usecaseCheckBlacklist usecaseBlacklistContract.ICheckBlacklist,
	// 	usecaseRemoveBlacklist usecaseBlacklistContract.IRemoveBlackList,
	// 	register_audit contracts_infrastructure.IAuditLogger,
	// 	persistence_cache contracts_infrastructure.ICache,
	// 	producer *producer.BlacklistProducer,
	// ) *service.BlacklistService {
	// 	return service.NewBlackListService(
	// 		usecaseCreateBlacklist,
	// 		usecaseCheckBlacklist,
	// 		usecaseRemoveBlacklist,
	// 		register_audit,
	// 		persistence_cache,
	// 		producer,
	// 	)
	// })

	return &ContainerInjection{c}
}

func (container *ContainerInjection) GetUsecaseAddBlacklist() (usecaseBlacklistContract.IAddBlacklist, error) {
	var useCase usecaseBlacklistContract.IAddBlacklist
	err := container.Invoke(func(u usecaseBlacklistContract.IAddBlacklist) {
		useCase = u
	})
	return useCase, err
}

func (container *ContainerInjection) GetUsecaseCheckBlacklist() (usecaseBlacklistContract.ICheckBlacklist, error) {
	var useCase usecaseBlacklistContract.ICheckBlacklist
	err := container.Invoke(func(u usecaseBlacklistContract.ICheckBlacklist) {
		useCase = u
	})
	return useCase, err
}

func (container *ContainerInjection) GetUsecaseRemoveBlacklist() (usecaseBlacklistContract.IRemoveBlackList, error) {
	var useCase usecaseBlacklistContract.IRemoveBlackList
	err := container.Invoke(func(u usecaseBlacklistContract.IRemoveBlackList) {
		useCase = u
	})
	return useCase, err
}

func (container *ContainerInjection) GetDispatcher() (*queue.Dispatcher, error) {
	var dispatcher *queue.Dispatcher
	err := container.Invoke(func(d *queue.Dispatcher) {
		dispatcher = d
	})
	return dispatcher, err
}

func (container *ContainerInjection) GetBlacklistService() (*service.BlacklistService, error) {
	var blacklistService *service.BlacklistService
	err := container.Invoke(func(s *service.BlacklistService) {
		blacklistService = s
	})
	return blacklistService, err
}
