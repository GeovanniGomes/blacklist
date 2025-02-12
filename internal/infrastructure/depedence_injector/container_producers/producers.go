package container_producers

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/producer"
	"go.uber.org/dig"
)

func RegisterProducers(c *dig.Container) {

	c.Provide(func(dispatcher *queue.Dispatcher) *producer.BlacklistProducer {
		return producer.NewBlacklistProducer(dispatcher)
	})
}
