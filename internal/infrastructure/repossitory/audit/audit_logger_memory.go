package audit

import (
	"encoding/json"
	"log"
)

type AuditLoggerMemory struct {

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

	l.collection = append(l.collection, detailsJSON)
	return err
}
