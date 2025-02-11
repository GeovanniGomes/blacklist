package container_usecase

import (
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	usecaseBlacklistContract "github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	usecseApplicationBlacklist "github.com/GeovanniGomes/blacklist/internal/application/usecase"
	"go.uber.org/dig"
)

func RegisterUseCase(c *dig.Container) {
	c.Provide(func(repository_blacklist repository.IBlackListRepository) usecaseBlacklistContract.IAddBlacklist {
		return usecseApplicationBlacklist.NewAddBlacklist(repository_blacklist)
	})

	c.Provide(func(repository_blacklist repository.IBlackListRepository) usecaseBlacklistContract.ICheckBlacklist {
		return usecseApplicationBlacklist.NewCheckBlacklist(repository_blacklist)
	})

	c.Provide(func(repository_blacklist repository.IBlackListRepository) usecaseBlacklistContract.IRemoveBlackList {
		return usecseApplicationBlacklist.NewRemoveBlacklist(repository_blacklist)
	})
}
