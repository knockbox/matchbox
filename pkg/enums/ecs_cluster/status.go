package ecs_cluster

type Status string

const (
	Active         Status = "active"
	Provisioning          = "provisioning"
	Deprovisioning        = "deprovisioning"
	Failed                = "failed"
	Inactive              = "inactive"
)
