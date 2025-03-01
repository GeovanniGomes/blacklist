package contracts

const (
	CHECK_BLACKLIST = "check"
	ADD_BLACKLIST   = "add blacklist"
)

type IAuditLogger interface {
	LogAction(userIdentifier int, action string, eventId *string ,details map[string]interface{}) error
}
