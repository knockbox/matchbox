package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type VPCInstanceAccessor interface {
	Create(vpc models.VPCInstance) (sql.Result, error)
	GetByDeploymentId(id int) (*models.VPCInstance, error)
}
