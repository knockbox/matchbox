package models

import (
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums/event"
	"github.com/knockbox/matchbox/pkg/payloads"
)

// EventParticipant represents a User participating in an Event.
type EventParticipant struct {
	Id            uint         `db:"id"`
	EventId       uint         `db:"event_id"`
	ParticipantId uuid.UUID    `db:"participant_id"`
	TeamId        uuid.UUID    `db:"team_id"`
	Status        event.Status `db:"status"`
	CanInvite     bool         `db:"can_invite"`
	CanManage     bool         `db:"can_manage"`
}

func NewEventParticipant(ev *Event, id uuid.UUID) *EventParticipant {
	return &EventParticipant{
		Id:            0,
		EventId:       ev.Id,
		ParticipantId: id,
		Status:        event.Member,
		CanInvite:     false,
		CanManage:     false,
	}
}

func (p *EventParticipant) ApplyCreate(payload *payloads.EventParticipantCreate) {
	p.Status = payload.Status
	p.CanInvite = *payload.CanInvite
	p.CanManage = *payload.CanManage
}

func (p *EventParticipant) DTO() *EventParticipantDTO {
	return &EventParticipantDTO{
		ParticipantId: p.ParticipantId,
		Status:        p.Status,
		CanInvite:     p.CanInvite,
		CanManage:     p.CanManage,
	}
}

type EventParticipantDTO struct {
	ParticipantId uuid.UUID    `json:"participant_id"`
	Status        event.Status `json:"status"`
	CanInvite     bool         `json:"can_invite"`
	CanManage     bool         `json:"can_manage"`
}
