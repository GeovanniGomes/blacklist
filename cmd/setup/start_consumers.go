package setup

import (
	"log"
	"os"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/consumer"
)

func StartQueueConsumers(container depedence_injector.ContainerInjection) {
	brokenHandler, err := container.GetNewRabbitMQ()
	if err != nil {
		log.Printf("Erro instance rabbitmq %v", err)
	}
	blacklistNotifyConsumer := consumer.NewBlacklistConsumer(brokenHandler)
	blacklistConsumerReport, err := container.GetBlacklistReportConsumer()

	if err != nil {
		println("DEU RUIM")
	}

	consumers := map[string]func([]byte) error{
		os.Getenv("QUEUE_BLACKLIST"):        blacklistNotifyConsumer.HandleMessage(), 
		os.Getenv("QUEUE_REPORT_BLACKLIST"): blacklistConsumerReport.HandleMessage(), 
	}

	for queueName, handler := range consumers {
		go func(queue string, h func([]byte) error) {
			err := brokenHandler.Consume(queue, func(msg []byte) {
				if err := h(msg); err != nil {
					log.Printf("❌ Erro ao processar mensagem da fila %s: %v", queue, err)
				}
			})
			if err != nil {
				log.Printf("❌ Erro ao iniciar consumidor %s: %v", queue, err)
			}
		}(queueName, handler)
	}

	log.Println("🎯 Todos os consumidores foram iniciados.")
	//consumerLoader := queue.NewConsumerLoader(brokenHandler)
	//consumerLoader.StartConsumers()

}
