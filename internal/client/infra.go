package client

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/matchbox/internal/platform"
	"github.com/knockbox/matchbox/pkg/accessors"
	deployment2 "github.com/knockbox/matchbox/pkg/enums/deployment"
	"github.com/knockbox/matchbox/pkg/models"
	"github.com/knockbox/matchbox/pkg/payloads"
)

type Infra struct {
	amz *Amazon
	dep accessors.DeploymentAccessor
}

func NewInfra(db *sqlx.DB, l hclog.Logger) *Infra {
	return &Infra{
		amz: NewAmazon(db, l),
		dep: platform.DeploymentSQLImpl{
			DB: db,
		},
	}
}

func (i *Infra) CreateDeployment(event *models.Event) error {
	deployment := models.NewDeployment(event)
	result, err := i.dep.Create(*deployment)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	deployment.Id = uint(id)

	err = i.amz.InitForDeployment(int(id))
	if err != nil {
		return err
	}

	_, _ = i.dep.UpdateStatusById(int(id), deployment2.Idle)

	return nil
}

func (i *Infra) GetDeploymentForEvent(event *models.Event) (*models.Deployment, error) {
	deployment, err := i.dep.GetDeploymentByActivityId(event.ActivityId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return deployment, err
}

func (i *Infra) CreateTaskDefinitionForEvent(event *models.Event, payload *payloads.TaskDefinitionCreatePayload) error {
	// Ensure the deployment exists.
	dep, err := i.GetDeploymentForEvent(event)
	if err != nil {
		return err
	}
	if dep == nil {
		return ErrDeploymentDoesNotExist
	}
	if dep.Status != deployment2.Idle {
		return ErrDeploymentNotReady
	}

	_, err = i.amz.CreateTaskDefinition(dep, payload)
	return err
}

func (i *Infra) GetTaskDefinitionForEvent(event *models.Event) (*models.ECSTaskDefinition, error) {
	// Ensure the deployment exists.
	dep, err := i.GetDeploymentForEvent(event)
	if err != nil {
		return nil, err
	}
	if dep == nil {
		return nil, ErrDeploymentDoesNotExist
	}
	if dep.Status != deployment2.Idle {
		return nil, ErrDeploymentNotReady
	}

	def, err := i.amz.GetTaskDefinition(int(dep.Id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return def, err
}

func (i *Infra) StartTaskForEvent(event *models.Event, flags []models.EventFlag, owner uuid.UUID) error {
	// Ensure the deployment exists.
	dep, err := i.GetDeploymentForEvent(event)
	if err != nil {
		return err
	}
	if dep == nil {
		return ErrDeploymentDoesNotExist
	}
	if dep.Status != deployment2.Idle {
		return ErrDeploymentNotReady
	}

	_, err = i.amz.StartTask(dep, owner, flags)
	return err
}

func (i *Infra) GetTaskForEvent(event *models.Event, owner uuid.UUID) (*models.ECSTaskInstance, error) {
	// Ensure the definition exists.
	def, err := i.GetTaskDefinitionForEvent(event)
	if err != nil {
		return nil, err
	}
	if def == nil {
		return nil, ErrTaskDefDoesNotExist
	}

	return i.amz.GetAndUpdateTask(int(def.Id), owner)
}

func (i *Infra) StopTaskForEvent(event *models.Event, owner uuid.UUID) error {
	// Ensure the definition exists.
	def, err := i.GetTaskDefinitionForEvent(event)
	if err != nil {
		return err
	}
	if def == nil {
		return ErrTaskDefDoesNotExist
	}

	return i.amz.StopTask(int(def.Id), owner)
}
