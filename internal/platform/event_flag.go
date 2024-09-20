package platform

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventFlagSQLImpl struct {
	*sqlx.DB
}

func (s EventFlagSQLImpl) Create(flag models.EventFlag) (sql.Result, error) {
	return utils.Transact(s.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertEventFlag, flag.EventId, flag.FlagId, flag.Difficulty, flag.EnvVar)
	})
}

func (s EventFlagSQLImpl) Update(flag models.EventFlag) (sql.Result, error) {
	return utils.Transact(s.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.UpdateEventFlag, flag.Difficulty, flag.EnvVar, flag.FlagId)
	})
}

func (s EventFlagSQLImpl) GetAllForEvent(id int) ([]models.EventFlag, error) {
	var flags []models.EventFlag
	err := s.Select(&flags, queries.SelectAllEventFlags, id)
	return flags, err
}

func (s EventFlagSQLImpl) DeleteByFlagId(flagId uuid.UUID) (sql.Result, error) {
	return utils.Transact(s.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.DeleteEventFlag, flagId)
	})
}
