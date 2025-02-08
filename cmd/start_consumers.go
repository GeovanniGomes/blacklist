package main

import (
	"os"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"
)

func StartQueueConsumers() {
	brokenQueue := queue.NewRabbitMQQueue(os.Getenv("CONNECTION_STRING_BROKEN_QUEUE"))
	consumerLoader := queue.NewConsumerLoader(brokenQueue)
	consumerLoader.StartConsumers()

}
