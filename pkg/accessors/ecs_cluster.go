package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type ECSClusterAccessor interface {
	Create(cluster models.ECSCluster) (sql.Result, error)
}
