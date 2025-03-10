package interfaces

type BlacklistInput struct {
	EventId        string          `json:"event_id" binding:"required"`
	Reason         string          `json:"reason" binding:"required"`
	Document       string          `json:"document" binding:"required"`
	Scope          string          `json:"scope" binding:"required"`
	UserIdentifier int             `json:"user_identifier" binding:"required"`
	BlockedUntil   *CustomDateTime `json:"blocked_until"`
}

type BlacklistOutputCheck struct {
	IsBlocked bool   `json:"is_blocked"`
	Reason    string `json:"reason"`
}

type BlacklistInputCheck struct {
	UserIdentifier int    `json:"user_identifier" binding:"required"`
	EventId        string `json:"event_id" binding:"required"`
}
type BlacklistInputRemove struct {
	UserIdentifier int    `json:"user_identifier" binding:"required"`
	EventId        string `json:"event_id" binding:"required"`
}

type BlacklistInputReport struct {
	StartDate CustomDate `json:"start_date" binding:"required"`
	EndDate   CustomDate `json:"end_date" binding:"required"`
}
