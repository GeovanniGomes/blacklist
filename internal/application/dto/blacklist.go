package dto

import "time"

type BlacklistInput struct {
	EventId        string     `json:"event_id"`
	Reason         string     `json:"reason"`
	Document       string     `json:"document"`
	Scope          string     `json:"scope"`
	UserIdentifier int        `json:"user_identifier"`
	BlockedUntil   *time.Time `json:"blocled_until"`
}

type BlacklistOutputCheck struct {
	IsBlocked bool   `json:"is_blocked"`
	Reason    string `json:"reason"`
}

type BlacklistInputCheck struct {
	UserIdentifier int    `json:"user_identifier"`
	EventId        string `json:"event_id"`
}
type BlacklistInputRemove struct {
	UserIdentifier int    `json:"user_identifier"`
	EventId        string `json:"event_id"`
}
