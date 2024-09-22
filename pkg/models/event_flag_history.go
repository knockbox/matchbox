package models

import (
	"github.com/google/uuid"
	"time"
)

// EventFlagHistory contains the history for an EventFlag within a given Event.
type EventFlagHistory struct {
	Id         uint      `db:"id"`
	EventId    uint      `db:"event_id"`
	FlagId     uint      `db:"flag_id"`
	Timestamp  time.Time `db:"timestamp"`
	RedeemerId uuid.UUID `db:"redeemer_id"`
}

func NewFlagHistory(event *Event, participant *EventParticipant, flag *EventFlag) *EventFlagHistory {
	return &EventFlagHistory{
		Id:         0,
		EventId:    event.Id,
		FlagId:     flag.Id,
		Timestamp:  time.Now(),
		RedeemerId: participant.ParticipantId,
	}
}
