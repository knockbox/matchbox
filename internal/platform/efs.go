package platform

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type EFSInstanceSQLImpl struct {
	*sqlx.DB
}

func (e EFSInstanceSQLImpl) GetByDeploymentId(id int) (*models.EFSInstance, error) {
	efs := &models.EFSInstance{}
	err := e.Get(efs, queries.SelectEFS, id)
	return efs, err
}

func (e EFSInstanceSQLImpl) Create(efsi models.EFSInstance) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertEFS, efsi.DeploymentId, efsi.AWSFileSystemId, efsi.AwsResourceId, efsi.State)
	})
}
