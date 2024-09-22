package platform

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type EventParticipantSQLImpl struct {
	*sqlx.DB
}

func (e EventParticipantSQLImpl) Create(participant models.EventParticipant) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertParticipant, participant.EventId, participant.ParticipantId, participant.Status, participant.CanInvite, participant.CanManage)
	})
}

func (e EventParticipantSQLImpl) Update(participant models.EventParticipant) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.UpdateParticipant, participant.Status, participant.CanInvite, participant.CanManage, participant.ParticipantId)
	})
}

func (e EventParticipantSQLImpl) GetAllByEventId(id int) ([]models.EventParticipant, error) {
	var participants []models.EventParticipant
	err := e.Select(&participants, queries.SelectAllParticipants, id)
	return participants, err
}

func (e EventParticipantSQLImpl) GetByEventAndParticipantId(eventId int, participantId uuid.UUID) (*models.EventParticipant, error) {
	participant := &models.EventParticipant{}
	err := e.Get(participant, queries.SelectParticipantByEventAndId, participantId, eventId)
	return participant, err
}
