package models

import (
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums/event"
)

// EventParticipant represents a User participating in an Event.
type EventParticipant struct {
	Id            uint         `db:"id"`
	EventId       uint         `db:"event_id"`
	ParticipantId uuid.UUID    `db:"participant_id"`
	Status        event.Status `db:"status"`
	CanInvite     bool         `db:"can_invite"`
	CanManage     bool         `db:"can_manage"`
}
