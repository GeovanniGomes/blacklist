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

func (c *BlacklistConsumer) HandleMessage() func([]byte) error {
	return func(message []byte) error {
		log.Printf("Process menssage blacklist: %s", message)
		return nil
	}
}
