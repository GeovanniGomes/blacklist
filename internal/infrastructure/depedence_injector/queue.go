package depedence_injector

import (
	"os"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/producer"
	"go.uber.org/dig"
)

func RegisterQueue(c *dig.Container) {
	c.Provide(func() contracts.IQueue {
		return queue.NewRabbitMQQueue(os.Getenv("CONNECTION_STRING_BROKEN_QUEUE"))
	})

	c.Provide(func(handler contracts.IQueue) *queue.Dispatcher {
		return queue.NewDispatcher(handler)
	})

	c.Provide(func(dispatcher *queue.Dispatcher) *producer.BlacklistProducer {
		return producer.NewBlacklistProducer(dispatcher)
	})
}