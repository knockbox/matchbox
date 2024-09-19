package platform

import (
	"database/sql"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventSQLImpl struct {
	*sqlx.DB
	hclog.Logger
}

func (e EventSQLImpl) Create(event models.Event) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertEvent, event.ActivityId, event.OrganizerId, event.Name, event.StartsAt, event.EndsAt, event.ImageName, event.ImageRepo, event.ImageTag, event.Private)
	})
}

func (e EventSQLImpl) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := e.Select(&events, queries.SelectAllEvents)
	return events, err
}

func (e EventSQLImpl) GetByActivityId(activityId string) (*models.Event, error) {
	event := &models.Event{}
	err := e.Get(event, queries.SelectEventByActivityId, activityId)
	return event, err
}
