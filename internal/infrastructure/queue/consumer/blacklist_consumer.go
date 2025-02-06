package consumer

import (
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

type BlacklistConsumer struct {
	queue contracts.IQueue
}

func NewBlacklistConsumer(queue contracts.IQueue) *BlacklistConsumer {
	return &BlacklistConsumer{queue: queue}
}

func (c *BlacklistConsumer) ProcessMessages() {
	err := c.queue.Consume("blacklist-queue", func(message []byte) {
		log.Printf("Processando mensagem: %s", message)
	})

	if err != nil {
		log.Fatalf("Erro ao consumir mensagens: %v", err)
	}
}
