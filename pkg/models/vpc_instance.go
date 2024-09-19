package models

import "github.com/knockbox/matchbox/pkg/enums/vpc_instance"

// VPCInstance represents a vpc on AWS.
type VPCInstance struct {
	Id            uint               `db:"id"`
	DeploymentId  uint               `db:"deployment_id"`
	AwsResourceId string             `db:"aws_resource_id"`
	State         vpc_instance.State `db:"state"`
}
