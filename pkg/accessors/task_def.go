package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type ECSTaskDefinitionAccessor interface {
	Create(def models.ECSTaskDefinition) (sql.Result, error)
	GetByDeploymentId(id int) (*models.ECSTaskDefinition, error)
}
