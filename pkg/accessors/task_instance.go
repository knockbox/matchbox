package accessors

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/models"
)

type TaskInstanceAccessor interface {
	Create(task models.ECSTaskInstance) (sql.Result, error)
	Select(taskDefId int, owner uuid.UUID) (*models.ECSTaskInstance, error)
	Update(task models.ECSTaskInstance) (sql.Result, error)
	Delete(taskDefId int, owner uuid.UUID) (sql.Result, error)
}
