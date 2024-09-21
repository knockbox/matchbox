package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventParticipantAccessor interface {
	Create(participant models.EventParticipant) (sql.Result, error)
	Update(participant models.EventParticipant) (sql.Result, error)
	GetAllByEventId(id int) ([]models.EventParticipant, error)
}
