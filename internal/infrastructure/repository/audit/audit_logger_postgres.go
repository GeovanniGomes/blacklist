package audit

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	uuid "github.com/satori/go.uuid"
)

var _ contracts.IAuditLogger = (*AuditLoggerRepository)(nil)

type AuditLoggerRepository struct {
	mutex       sync.Mutex
	persistence contracts.IDatabaseRelational
}

func NewDBAuditLogger(persistence contracts.IDatabaseRelational) *AuditLoggerRepository {
	return &AuditLoggerRepository{persistence: persistence}
}

func (a *AuditLoggerRepository) LogAction(userIdentifier int, eventId, action string, details map[string]interface{}) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	detailsJSON, err := json.Marshal(details)

	if err != nil {
		log.Printf("error converter detalhes para JSON: %v", err)
		return err
	}
	err = a.persistence.InsertData(
		"auditlog",
		[]string{"id", "user_identifier", "event_id", "action", "details"},
		[]interface{}{
			uuid.NewV4().String(),
			userIdentifier,
			eventId,
			action,
			detailsJSON,
		},
	)
	if err != nil {
		log.Fatalf("Error peersiste data auditlogger: %v", err)
		return err
	}
	return nil
}
