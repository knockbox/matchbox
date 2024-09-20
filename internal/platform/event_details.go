package platform

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventDetailsDQLImpl struct {
	*sqlx.DB
}

func (e EventDetailsDQLImpl) CreateForEvent(id int) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertEventDetails, id)
	})
}

func (e EventDetailsDQLImpl) Update(details models.EventDetails) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.UpdateEventDetails, details.ProfilePicture, details.Description, details.GithubURL, details.TwitterURL, details.WebsiteURL)
	})
}
