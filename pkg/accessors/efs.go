package accessors

import (
	"database/sql"
	"github.com/knockbox/matchbox/pkg/models"
)

type EFSInstanceAccessor interface {
	Create(efsi models.EFSInstance) (sql.Result, error)
}
