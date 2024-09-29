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

func (f *EventFlagHistory) DTO() *EventFlagHistoryDTO {
	return &EventFlagHistoryDTO{
		EventId:    f.EventId,
		FlagId:     f.FlagId,
		Timestamp:  f.Timestamp,
		RedeemerId: f.RedeemerId,
	}
}

type EventFlagHistoryDTO struct {
	EventId    uint      `json:"event_id"`
	FlagId     uint      `json:"flag_id"`
	Timestamp  time.Time `json:"timestamp"`
	RedeemerId uuid.UUID `json:"redeemer_id"`
}
