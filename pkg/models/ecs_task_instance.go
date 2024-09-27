package models

import (
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums/ecs_task_instance"
	"time"
)

// ECSTaskInstance represents a single instance for a given ECS Task.
type ECSTaskInstance struct {
	Id                  uint                     `db:"id"`
	AwsArn              string                   `db:"aws_arn"`
	ECSTaskDefinitionId uint                     `db:"ecs_task_definition_id"`
	ECSClusterId        uint                     `db:"ecs_cluster_id"`
	PullStart           *time.Time               `db:"pull_start"`
	PullStop            *time.Time               `db:"pull_stop"`
	StartedAt           *time.Time               `db:"started_at"`
	StoppedAt           *time.Time               `db:"stopped_at"`
	StoppedReason       *string                  `db:"stopped_reason"`
	Status              ecs_task_instance.Status `db:"status"`
	InstanceOwnerId     uuid.UUID                `db:"instance_owner_id"`
}

func NewTaskInstance(taskDefId, clusterId uint, owner uuid.UUID) *ECSTaskInstance {
	return &ECSTaskInstance{
		Id:                  0,
		AwsArn:              "",
		ECSTaskDefinitionId: taskDefId,
		ECSClusterId:        clusterId,
		PullStart:           nil,
		PullStop:            nil,
		StartedAt:           nil,
		StoppedAt:           nil,
		StoppedReason:       nil,
		Status:              ecs_task_instance.Unknown,
		InstanceOwnerId:     owner,
	}
}

// UpdateFromTask sets fields based on the given types.Task.
func (e *ECSTaskInstance) UpdateFromTask(task types.Task) {
	e.AwsArn = *task.TaskArn
	e.StartedAt = task.StartedAt
	e.StoppedAt = task.StoppedAt
	e.PullStart = task.PullStartedAt
	e.PullStop = task.PullStoppedAt
	e.StoppedReason = task.StoppedReason

	switch task.HealthStatus {
	case types.HealthStatusHealthy:
		e.Status = ecs_task_instance.Healthy
	case types.HealthStatusUnhealthy:
		e.Status = ecs_task_instance.Unhealthy
	case types.HealthStatusUnknown:
		e.Status = ecs_task_instance.Unknown
	}
}
