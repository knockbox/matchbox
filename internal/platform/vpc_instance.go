package platform

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/queries"
	"github.com/knockbox/matchbox/pkg/models"
)

type VPCInstanceSQLImpl struct {
	*sqlx.DB
}

func (v VPCInstanceSQLImpl) Create(vpc models.VPCInstance) (sql.Result, error) {
	return utils.Transact(v.DB, func(tx *sql.Tx) (sql.Result, error) {
		return tx.Exec(queries.InsertVPCInstance, vpc.DeploymentId, vpc.AwsResourceId, vpc.SubnetID, vpc.SecurityGroupID, vpc.InternetGatewayID, vpc.State)
	})
}
