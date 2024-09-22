package models

import "github.com/knockbox/matchbox/pkg/enums/vpc_instance"

// VPCInstance represents a vpc on AWS.
type VPCInstance struct {
	Id                uint               `db:"id"`
	DeploymentId      uint               `db:"deployment_id"`
	AwsResourceId     string             `db:"aws_resource_id"`
	SubnetID          string             `db:"subnet_id"`
	SecurityGroupID   string             `db:"security_group_id"`
	InternetGatewayID string             `db:"internet_gateway_id"`
	State             vpc_instance.State `db:"state"`
}

func NewVPCInstance(deploymentId int) *VPCInstance {
	return &VPCInstance{
		Id:                0,
		DeploymentId:      uint(deploymentId),
		AwsResourceId:     "",
		SubnetID:          "",
		SecurityGroupID:   "",
		InternetGatewayID: "",
		State:             vpc_instance.Pending,
	}
}
