package accessors

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventParticipantAccessor interface {
	Create(participant models.EventParticipant) (sql.Result, error)
	Update(participant models.EventParticipant) (sql.Result, error)
	GetAllByEventId(id int) ([]models.EventParticipant, error)
	GetByEventAndParticipantId(eventId int, participantId uuid.UUID) (*models.EventParticipant, error)
}
