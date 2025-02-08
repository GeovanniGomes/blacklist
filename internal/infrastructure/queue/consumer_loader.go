package queue

import (
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/consumer"
)

type ConsumerLoader struct {
	brokenHandler contracts.IQueue
}

func NewConsumerLoader(brokenHandler contracts.IQueue) *ConsumerLoader {
	return &ConsumerLoader{brokenHandler: brokenHandler}
}

func (c *ConsumerLoader) StartConsumers() {
	blacklistConsumer := consumer.NewBlacklistConsumer(c.brokenHandler)

	consumers := map[string]func([]byte) error{
		"blacklist-queue": blacklistConsumer.HandleMessage(),
	}

	for queueName, handler := range consumers {
		go func(queue string, h func([]byte) error) {
			err := c.brokenHandler.Consume(queue, func(msg []byte) {
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
}
