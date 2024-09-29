package platform

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventFlagHistorySQLImpl struct {
	*sqlx.DB
}

func (e EventFlagHistorySQLImpl) GetByEvent(eventId int) ([]models.EventFlagHistory, error) {
	var history []models.EventFlagHistory
	err := e.Select(&history, queries.SelectFlagHistoryByEvent, eventId)
	return history, err
}

func (e EventFlagHistorySQLImpl) Create(history models.EventFlagHistory) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertFlagHistory, history.EventId, history.FlagId, history.RedeemerId)
	})
}

func (e EventFlagHistorySQLImpl) GetByEventFlagRedeemer(eventId int, flagId int, redeemer uuid.UUID) (*models.EventFlagHistory, error) {
	history := &models.EventFlagHistory{}
	err := e.Get(history, queries.SelectFlagHistoryByRedeemer, eventId, flagId, redeemer)
	return history, err
}
