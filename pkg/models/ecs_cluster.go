package models

import "github.com/knockbox/matchbox/pkg/enums/ecs_cluster"

// ECSCluster represents an ECS cluster.
type ECSCluster struct {
	Id           uint               `db:"id"`
	AwsArn       string             `db:"aws_arn"`
	ClusterName  string             `db:"cluster_name"`
	DeploymentId uint               `db:"deployment_id"`
	Status       ecs_cluster.Status `db:"status"`
}
