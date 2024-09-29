package accessors

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventFlagHistoryAccessor interface {
	Create(history models.EventFlagHistory) (sql.Result, error)
	GetByEventFlagRedeemer(eventId int, flagId int, redeemer uuid.UUID) (*models.EventFlagHistory, error)
	GetByEvent(eventId int) ([]models.EventFlagHistory, error)
}
