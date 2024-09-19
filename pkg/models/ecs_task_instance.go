package models

import (
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
	PullStart           time.Time                `db:"pull_start"`
	PullEnd             time.Time                `db:"pull_end"`
	StartedAt           time.Time                `db:"started_at"`
	StoppedAt           time.Time                `db:"stopped_at"`
	StoppedReason       string                   `db:"stopped_reason"`
	Status              ecs_task_instance.Status `db:"status"`
	InstanceOwnerId     uuid.UUID                `db:"instance_owner_id"`
}
