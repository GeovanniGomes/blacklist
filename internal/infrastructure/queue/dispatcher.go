package queue

import (
	"encoding/json"
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

type Dispatcher struct {
	queue contracts.IQueue
}

func NewDispatcher(queue contracts.IQueue) *Dispatcher {
	return &Dispatcher{queue: queue}
}

func (d *Dispatcher) Dispatch(queue, eventType string, data interface{}) {
	// Criando uma estrutura que representa a mensagem
	messageStruct := map[string]interface{}{
		"event": eventType,
		"data":  data, // Agora `data` é tratado como JSON válido
	}

	// Serializando para JSON corretamente
	message, err := json.Marshal(messageStruct)
	if err != nil {
		log.Printf("Erro ao serializar mensagem: %v", err)
		return
	}

	// Publicando na fila
	err = d.queue.Publish(queue, message)
	if err != nil {
		log.Printf("Erro ao enviar mensagem: %v", err)
	}
}
