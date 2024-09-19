package ecs_task_instance

type Status string

const (
	Healthy   Status = "healthy"
	Unhealthy        = "unhealthy"
	Unknown          = "unknown"
)
