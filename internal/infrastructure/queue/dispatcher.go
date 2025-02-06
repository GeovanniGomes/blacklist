package queue

import (
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

type Dispatcher struct {
	queue contracts.QueueInterface
}

func NewDispatcher(queue contracts.QueueInterface) *Dispatcher {
	return &Dispatcher{queue: queue}
}

func (d *Dispatcher) Dispatch(queue,eventType, data string) {
	message := []byte(`{"event":"` + eventType + `","data":"` + data + `"}`)

	err := d.queue.Publish(queue, message)
	if err != nil {
		log.Printf("Erro ao enviar mensagem: %v", err)
	}
}
