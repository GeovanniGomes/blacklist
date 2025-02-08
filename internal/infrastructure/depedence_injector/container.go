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
