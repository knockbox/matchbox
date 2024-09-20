package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventDetailsAccessor interface {
	CreateForEvent(id int) (sql.Result, error)
	Update(details models.EventDetails) (sql.Result, error)
}
