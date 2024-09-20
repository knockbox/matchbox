package accessors

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventFlagAccessor interface {
	Create(flag models.EventFlag) (sql.Result, error)
	Update(flag models.EventFlag) (sql.Result, error)
	GetAllForEvent(id int) ([]models.EventFlag, error)
	DeleteByFlagId(flagId uuid.UUID) (sql.Result, error)
}
