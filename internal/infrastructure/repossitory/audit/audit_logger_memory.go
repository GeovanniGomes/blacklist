package audit

import (
	"encoding/json"
	"log"
	"sync"
)

type AuditLoggerMemory struct {
	mu         sync.Mutex
	collection [][]byte
}

func (l *AuditLoggerMemory) LogAction(userIdentifier int, eventId, action string, details *map[string]interface{}) error{
	logEntry := map[string]interface{}{
		"user_id": userIdentifier,
		"eventId": eventId,
		"action":  action,
		"details": details,
	}
	detailsJSON, err := json.Marshal(logEntry)
	if err != nil {
		log.Printf("error converter detalhes para JSON: %v", err)
		return err
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.collection = append(l.collection, detailsJSON)
	return err
}
