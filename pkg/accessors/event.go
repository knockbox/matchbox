package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventAccessor interface {
	Create(event models.Event) (sql.Result, error)
	GetAll() ([]models.Event, error)
	GetByActivityId(activityId string) (*models.Event, error)
}
