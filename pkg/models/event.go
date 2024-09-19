package models

import (
	"github.com/google/uuid"
	"time"
)

// Event represents an Event
type Event struct {
	Id          uint      `db:"id"`
	ActivityId  uuid.UUID `db:"activity_id"`
	OrganizerId uuid.UUID `db:"organizer_id"`
	Name        string    `db:"name"`
	StartsAt    time.Time `db:"starts_at"`
	EndsAt      time.Time `db:"ends_at"`
	ImageName   string    `db:"image_name"`
	ImageTag    string    `db:"image_tag"`
	Private     bool      `db:"private"`
}
