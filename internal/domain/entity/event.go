package entity

import (
	"errors"
	"time"

	"github.com/GeovanniGomes/blacklist/internal/domain/value_objects"
	"github.com/GeovanniGomes/blacklist/internal/util"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Event struct {
	id          string
	title       string
	description string
	date        time.Time
	category    value_objects.Category
	isActive    bool
	status      string
	createdAt   time.Time
}

func NewEvent(id, title, description string, date, createdAt time.Time, category value_objects.Category, isActive bool, status string) *Event {
	new_event := Event{
		id:          id,
		title:       title,
		description: description,
		date:        date,
		category:    category,
		isActive:    isActive,
		createdAt:   createdAt,
		status:      status,
	}

	return &new_event
}

func (event *Event) Enable() error {
	now := util.TruncateDate(time.Now())

	if now.After(util.TruncateDate(event.date)) {
		return errors.New("it is not possible to enable an event with a past date")
	}
	event.status = ENABLED
	return nil
}

func (event *Event) Disable() error {
	if time.Now().Before(event.date) {
		return errors.New("it is not possible to disable an event with a past date")
	}
	event.status = DISABLED
	return nil
}

func (event *Event) ChangeCatrgory(category value_objects.Category) {
	event.category = category
}
func (event *Event) ChangeDateEvent(dateEvent time.Time) error {
	event.date = dateEvent
	if err := event.IsValid(); err != nil {
		return err
	}

	return nil
}

func (event *Event) IsValid() error {
	if util.GetSizeString(event.title) == 0 {
		return errors.New("the title is required")
	}
	if util.GetSizeString(event.description) == 0 {
		return errors.New("the description is required")
	}
	now := util.TruncateDate(time.Now())
	if event.date.Before(now) {
		return errors.New("Event date cannot be less than the current date")
	}
	return nil
}
func (event *Event) GetId() string {
	return event.id
}
func (event *Event) GetTitle() string {
	return event.title
}
func (event *Event) GetDescription() string {
	return event.description
}
func (event *Event) GetDate() time.Time {
	return event.date
}
func (event *Event) GetCategory() *value_objects.Category {
	return &event.category
}
func (event *Event) GetIsActive() bool {
	return event.isActive
}
func (event *Event) GetStatus() string {
	return event.status
}
func (event *Event) GetCreatedAt() time.Time {
	return event.createdAt
}
