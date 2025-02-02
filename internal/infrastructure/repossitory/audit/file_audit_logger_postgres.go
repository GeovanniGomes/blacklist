package audit

// import (
// 	"encoding/json"
// 	"log"

// 	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence/contracts"
// )

// type AuditLogger struct {
// 	db contracts.DatabaseRelationalInterface
// }

// func NewDBAuditLogger(db contracts.DatabaseRelationalInterface) *AuditLogger {
// 	return &AuditLogger{db: db}
// }

// func (l *AuditLogger) LogAction(userIdentifier int, eventId, action string, details *map[string]string) {
// 	detailsJSON, err := json.Marshal(details)
// 	if err != nil {
// 		log.Printf("Erro ao converter detalhes para JSON: %v", err)
// 		return
// 	}

// 	query := "INSERT INTO audit_logs (user_identifier, event_id, action, details) VALUES ($1, $2, $3, $4)"
// 	err = l.db.ExecuteQueryWithTransaction(query, userIdentifier, eventId, action, string(detailsJSON))
// 	if err != nil {
// 		log.Printf("Erro ao inserir log no banco: %v", err)
// 	}
// }
