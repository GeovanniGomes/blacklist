package depedence_injector

import (
	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	usecaseBlacklistContract "github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	usecseApplicationBlacklist "github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"go.uber.org/dig"
)

func RegisterUsecase(c *dig.Container) {

	c.Provide(func(repository repositoty.IBlackListRepository) usecaseBlacklistContract.IAddBlacklist {
		return usecseApplicationBlacklist.NewAddBlacklist(repository)
	})

	c.Provide(func(repository repositoty.IBlackListRepository) usecaseBlacklistContract.ICheckBlacklist {
		return usecseApplicationBlacklist.NewCheckBlacklist(repository)
	})

	c.Provide(func(repository repositoty.IBlackListRepository) usecaseBlacklistContract.IRemoveBlackList {
		return usecseApplicationBlacklist.NewRemoveBlacklist(repository)
	})
}
