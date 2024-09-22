package models

import (
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums/deployment"
)

// Deployment represents a deployed event.
type Deployment struct {
	Id         uint              `db:"id"`
	InstanceId uuid.UUID         `db:"instance_id"`
	EventId    uuid.UUID         `db:"event_id"`
	Status     deployment.Status `db:"status"`
}

func NewDeployment(event *Event) *Deployment {
	return &Deployment{
		Id:         0,
		InstanceId: uuid.New(),
		EventId:    event.ActivityId,
		Status:     deployment.Preparing,
	}
}
