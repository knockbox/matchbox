package models

import "github.com/aws/aws-sdk-go-v2/service/efs/types"

type EFSInstance struct {
	Id              uint                 `db:"id"`
	DeploymentId    uint                 `db:"deployment_id"`
	AWSFileSystemId string               `db:"aws_file_system_id"`
	AwsResourceId   string               `db:"aws_resource_id"`
	State           types.LifeCycleState `db:"state"`
}

func NewEFSInstance(id int) *EFSInstance {
	return &EFSInstance{
		Id:              0,
		DeploymentId:    uint(id),
		AWSFileSystemId: "",
		AwsResourceId:   "",
		State:           types.LifeCycleStateCreating,
	}
}
