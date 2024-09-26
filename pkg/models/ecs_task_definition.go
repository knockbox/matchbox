package models

import "github.com/google/uuid"

// ECSTaskDefinition represents a Task Definition.
type ECSTaskDefinition struct {
	Id           uint      `db:"id"`
	DeploymentId uint      `db:"deployment_id"`
	FamilyId     uuid.UUID `db:"family_id"`
	AwsArn       string    `db:"aws_arn"`
}

func NewECSTaskDefinition(deploymentId uint) *ECSTaskDefinition {
	return &ECSTaskDefinition{
		Id:           0,
		DeploymentId: deploymentId,
		FamilyId:     uuid.New(),
		AwsArn:       "",
	}
}
