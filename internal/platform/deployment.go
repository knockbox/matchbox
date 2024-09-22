package platform

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	deployment2 "github.com/knockbox/matchbox/pkg/enums/deployment"
	"github.com/knockbox/matchbox/pkg/models"
)

type DeploymentSQLImpl struct {
	*sqlx.DB
}

func (d DeploymentSQLImpl) Create(deployment models.Deployment) (sql.Result, error) {
	return utils.Transact(d.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertDeployment, deployment.InstanceId, deployment.EventId)
	})
}

func (d DeploymentSQLImpl) GetDeploymentByActivityId(id uuid.UUID) (*models.Deployment, error) {
	deployment := &models.Deployment{}
	err := d.Get(deployment, queries.SelectDeploymentByEventId, id)
	return deployment, err
}

func (d DeploymentSQLImpl) UpdateStatusById(id int, status deployment2.Status) (sql.Result, error) {
	return utils.Transact(d.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.UpdateDeploymentStatusById, status, id)
	})
}
