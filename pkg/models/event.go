package models

import (
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/payloads"
	"github.com/knockbox/matchbox/pkg/utils"
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
	ImageRepo   string    `db:"image_repo"`
	ImageTag    string    `db:"image_tag"`
	Private     bool      `db:"private"`
}

// NewEvent creates a new event with the ActivityId populated and the OrganizerId set to the provided uuid.
func NewEvent(organizer uuid.UUID) *Event {
	return &Event{
		Id:          0,
		ActivityId:  uuid.New(),
		OrganizerId: organizer,
		Name:        "",
		StartsAt:    time.Time{},
		EndsAt:      time.Time{},
		ImageName:   "",
		ImageRepo:   "",
		ImageTag:    "",
		Private:     false,
	}
}

// ApplyCreate ensures a valid event can be created and applies the values.
func (e *Event) ApplyCreate(payload *payloads.EventCreate) error {
	e.Name = payload.Name

	start, end, err := utils.ParseAndValidateTime(payload.StartsAt, payload.EndsAt)
	if err != nil {
		return err
	}

	e.StartsAt = start.UTC()
	e.EndsAt = end.UTC()

	e.ImageName = payload.ImageNamespace
	e.ImageRepo = payload.ImageRepository
	e.ImageTag = payload.ImageTag

	e.Private = *payload.Private

	return nil
}

// DTO converts the Event to the EventDTO.
func (e *Event) DTO() *EventDTO {
	return &EventDTO{
		Id:          nil,
		ActivityId:  e.ActivityId,
		OrganizerId: e.OrganizerId,
		Name:        e.Name,
		StartsAt:    e.StartsAt,
		EndsAt:      e.EndsAt,
		ImageName:   e.ImageName,
		ImageRepo:   e.ImageRepo,
		ImageTag:    e.ImageTag,
		Private:     e.Private,
	}
}

// EventDTO is used when returning an Event as JSON.
type EventDTO struct {
	Id          *uint     `json:"id,omitempty"`
	ActivityId  uuid.UUID `json:"activity_id"`
	OrganizerId uuid.UUID `json:"organizer_id"`
	Name        string    `json:"name"`
	StartsAt    time.Time `json:"starts_at"`
	EndsAt      time.Time `json:"ends_at"`
	ImageName   string    `json:"image_name"`
	ImageRepo   string    `json:"image_repo"`
	ImageTag    string    `json:"image_tag"`
	Private     bool      `json:"private"`
}
