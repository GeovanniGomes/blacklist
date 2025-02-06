package producer

import (
	"encoding/json"
	"log"
	"os"

	"github.com/GeovanniGomes/blacklist/internal/domain/entity"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"
)

type BlacklistProducer struct {
	dispatcher *queue.Dispatcher
}

func NewBlacklistProducer(dispatcher *queue.Dispatcher) *BlacklistProducer {
	return &BlacklistProducer{dispatcher: dispatcher}
}

func (p *BlacklistProducer) NotifyBlacklist(blacklist entity.BlackList) {
	detailMessage := map[string]interface{}{
		"id":              blacklist.GetId(),
		"user_identifier": blacklist.GetUserIdentifier(),
		"event_id":        blacklist.GetEventId(),
		"scope":           blacklist.GetScope(),
		"blocked_until":   blacklist.GetBlockedUntil(),
		"blocked_type":    blacklist.GetBlockedType(),
		"created_at":      blacklist.GetCreatedAt(),
	}

	message, _ := json.Marshal(detailMessage)
	p.dispatcher.Dispatch(os.Getenv("QUEUE_BLACKLIST"), "blacklist.created", string(message))

	log.Printf("Mensagem publicada na fila: %s", message)
}
