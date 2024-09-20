package models

import (
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums/difficulty"
	"github.com/knockbox/matchbox/pkg/payloads"
)

// EventFlag represents a generated for an Event.
type EventFlag struct {
	Id         uint                  `db:"id"`
	EventId    uint                  `db:"event_id"`
	FlagId     uuid.UUID             `db:"flag_id"`
	Difficulty difficulty.Difficulty `db:"difficulty"`
	EnvVar     string                `db:"env_var"`
}

func NewEventFlag(eventId uint) *EventFlag {
	return &EventFlag{
		Id:         0,
		EventId:    eventId,
		FlagId:     uuid.New(),
		Difficulty: "",
		EnvVar:     "",
	}
}

func (f *EventFlag) ApplyCreate(payload *payloads.EventFlagCreate) {
	f.Difficulty = payload.Difficulty
	f.EnvVar = payload.EnvVar
}

func (f *EventFlag) ApplyUpdate(payload *payloads.EventFlagUpdate) {
	if payload.Difficulty != nil {
		f.Difficulty = *payload.Difficulty
	}

	if payload.EnvVar != nil {
		f.EnvVar = *payload.EnvVar
	}
}

func (f *EventFlag) DTO() *EventFlagDTO {
	return &EventFlagDTO{
		FlagId:     f.FlagId,
		Difficulty: f.Difficulty,
		EnvVar:     f.EnvVar,
	}
}

type EventFlagDTO struct {
	FlagId     uuid.UUID             `json:"flag_id"`
	Difficulty difficulty.Difficulty `json:"difficulty"`
	EnvVar     string                `json:"env_var"`
}
