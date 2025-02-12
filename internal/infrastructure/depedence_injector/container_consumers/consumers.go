package container_consumers

import (
	repositoty "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/consumer"
	"go.uber.org/dig"
)

func RegisterConsumers(c *dig.Container) {

	c.Provide(func(handler contracts.IQueue) *consumer.BlacklistConsumer {
		return consumer.NewBlacklistConsumer(handler)
	})

	c.Provide(func(
		handler contracts.IQueue,
		clientUpload contracts.IFileSystem,
		repositotyBlacklist repositoty.IBlackListRepository,

	) *consumer.BlacklistReportConsumer {
		return consumer.NewBlacklistReportConsumer(handler, repositotyBlacklist, clientUpload)
	})
}
