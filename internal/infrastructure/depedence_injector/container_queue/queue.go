package container_queue

import (
	"log"
	"os"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"
	"go.uber.org/dig"
)

func RegisterBroken(c *dig.Container) {

	c.Provide(func() contracts.IQueue {
		connectionString := os.Getenv("CONNECTION_STRING_BROKEN_QUEUE")
		if connectionString == "" {
			log.Fatal("A variável de ambiente CONNECTION_STRING_BROKEN_QUEUE não está configurada.")
		}
		return queue.NewRabbitMQQueue(connectionString)
	})
	c.Provide(func(handler contracts.IQueue) *queue.Dispatcher {
		return queue.NewDispatcher(handler)
	})
}
