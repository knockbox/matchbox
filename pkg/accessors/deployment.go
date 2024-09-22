package accessors

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums/deployment"
	"github.com/knockbox/matchbox/pkg/models"
)

type DeploymentAccessor interface {
	Create(deployment models.Deployment) (sql.Result, error)
	GetDeploymentByActivityId(id uuid.UUID) (*models.Deployment, error)
	UpdateStatusById(id int, status deployment.Status) (sql.Result, error)
}
