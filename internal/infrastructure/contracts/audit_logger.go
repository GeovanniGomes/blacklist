package contracts

const (
	CHECK_BLACKLIST = "check"
	ADD_BLACKLIST   = "add blacklist"
)

type IAuditLogger interface {
	LogAction(userIdentifier int, eventId, action string, details map[string]interface{}) error
}
