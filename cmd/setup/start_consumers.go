package setup

import (
	"log"
	"os"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
)

func StartQueueConsumers(container depedence_injector.ContainerInjection) {
	brokenHandler, err := container.GetNewRabbitMQ()
	if err != nil {
		log.Printf("Error instance rabbitmq %v", err)
	}
	blacklistNotifyConsumer, _ := container.GetBlacklistConsumer()
	blacklistConsumerReport, _ := container.GetBlacklistReportConsumer()

	consumers := map[string]func([]byte) error{
		os.Getenv("QUEUE_BLACKLIST"):        blacklistNotifyConsumer.HandleMessage(),
		os.Getenv("QUEUE_REPORT_BLACKLIST"): blacklistConsumerReport.HandleMessage(),
	}

	for queueName, handler := range consumers {
		go func(queue string, h func([]byte) error) {
			err := brokenHandler.Consume(queue, func(msg []byte) {
				if err := h(msg); err != nil {
					log.Printf("❌ Error process message of queue %s: %v", queue, err)
				}
			})
			if err != nil {
				log.Printf("❌ error start consumer %s: %v", queue, err)
			}
		}(queueName, handler)
	}

	log.Println("🎯 All consumers have been initiated.")
}
