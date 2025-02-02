package audit

import (
	"log"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
)

var _ contracts.AuditLoggerInterface = (*AuditLoggerRepository)(nil)
type AuditLoggerRepository struct {
	persistence contracts.DatabaseRelationalInterface
}

func NewDBAuditLogger(persistence contracts.DatabaseRelationalInterface) *AuditLoggerRepository {
	return &AuditLoggerRepository{persistence: persistence}
}

func (l *AuditLoggerRepository) LogAction(userIdentifier int, eventId, action string, details *map[string]interface{}) error {
	err := l.persistence.InsertData(
		"blacklist",
		[]string{"id", "user_identifier", "event_id", "action", "details"},
		[]interface{}{
			userIdentifier,
			eventId,
			action,
			details,
		},
	)
	if err != nil {
		log.Fatalf("Error peersiste data auditlogger: %v", err)
		return err
	}
	return nil
}
