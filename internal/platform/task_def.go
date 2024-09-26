package platform

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type ECSTaskDefinitionSQLImpl struct {
	*sqlx.DB
}

func (e ECSTaskDefinitionSQLImpl) Create(def models.ECSTaskDefinition) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertTaskDef, def.DeploymentId, def.FamilyId, def.AwsArn)
	})
}

func (e ECSTaskDefinitionSQLImpl) GetByDeploymentId(id int) (*models.ECSTaskDefinition, error) {
	def := &models.ECSTaskDefinition{}
	err := e.Get(def, queries.SelectTaskDefByDeploymentId, id)
	return def, err
}
