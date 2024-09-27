package platform

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type ECSTaskInstanceSQLImpl struct {
	*sqlx.DB
}

func (e ECSTaskInstanceSQLImpl) Create(task models.ECSTaskInstance) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertTaskInstance, task.AwsArn, task.ECSTaskDefinitionId, task.ECSClusterId, task.InstanceOwnerId)
	})
}

func (e ECSTaskInstanceSQLImpl) Select(taskDefId int, owner uuid.UUID) (*models.ECSTaskInstance, error) {
	task := &models.ECSTaskInstance{}
	err := e.Get(task, queries.SelectTaskInstance, taskDefId, owner)
	return task, err
}

func (e ECSTaskInstanceSQLImpl) Update(task models.ECSTaskInstance) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.UpdateTaskInstance, task.PullStart, task.PullEnd, task.StartedAt, task.StoppedAt, task.StoppedReason, task.Status, task.ECSTaskDefinitionId, task.InstanceOwnerId)
	})
}

func (e ECSTaskInstanceSQLImpl) Delete(taskDefId int, owner uuid.UUID) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.DeleteTaskInstance, taskDefId, owner)
	})
}
