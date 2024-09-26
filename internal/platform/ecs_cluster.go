package platform

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type ECSClusterSQLImpl struct {
	*sqlx.DB
}

func (e ECSClusterSQLImpl) GetByDeploymentId(id int) (*models.ECSCluster, error) {
	cluster := &models.ECSCluster{}
	err := e.Get(cluster, queries.SelectCluster, id)
	return cluster, err
}

func (e ECSClusterSQLImpl) Create(cluster models.ECSCluster) (sql.Result, error) {
	return utils.Transact(e.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertCluster, cluster.AwsArn, cluster.ClusterName, cluster.DeploymentId, cluster.Status)
	})
}
