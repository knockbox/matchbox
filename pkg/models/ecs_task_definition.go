package models

import "github.com/google/uuid"

// ECSTaskDefinition represents a Task Definition.
type ECSTaskDefinition struct {
	Id       uint      `db:"id"`
	FamilyId uuid.UUID `db:"family_id"`
	AwsArn   string    `db:"aws_arn"`
}
