package payloads

// TaskDefinitionCreatePayload defines the payload required to register a task definition.
type TaskDefinitionCreatePayload struct {
	Containers []TaskContainerDefinition `json:"containers" validate:"required,gt=0,dive"`
	Volumes    []*TaskVolume             `json:"volumes" validate:"dive,omitempty,gte=0"`
	CPU        string                    `json:"cpu" validate:"required,numeric"`
	Memory     string                    `json:"memory" validate:"required,numeric"`
}

// TaskContainerDefinition defines the containers present in the task definition
type TaskContainerDefinition struct {
	// EnvironmentVars are the variables presented to the container. (these can be overridden when starting the task)
	EnvironmentVars []*ContainerVariable `json:"env" validate:"dive,omitempty,gte=0"`

	// Ports are the port mappings present in the container
	Ports []ContainerPortMapping `json:"ports" validate:"required,gt=0,dive"`

	// Volumes contains any additional volumes that can be attached to the container
	Volumes []*ContainerVolume `json:"volumes" validate:"dive,omitempty,gte=0"`

	// Image is the image name of the container: e.g. cesoun/knockbox:go-1.18.3
	Image string `json:"image" validate:"required"`

	// Essential determines if this container is required for the rest of the task to function.
	Essential *bool `json:"essential" validate:"omitempty"`
}

// ContainerVariable defines the environment variables set on the container
type ContainerVariable struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// ContainerPortMapping define the containers port mappings
type ContainerPortMapping struct {
	// ContainerPort defines the port to bind in the container
	ContainerPort int32 `json:"container_port" validate:"required"`

	// HostPort defines the port to bind to the host (if omitted, the ContainerPort will be used)
	HostPort *int32 `json:"host_port" validate:"omitempty"`
	Name     string `json:"name" validate:"required"`

	// Protocol is one of grpc, http, http2, or none (defaults to http)
	Protocol *string `json:"protocol" validate:"omitempty"`
}

// ContainerVolume defines the volumes mounted to the container
type ContainerVolume struct {
	// Path is where the volume is mounted. e.g. /mnt/efs
	Path string `json:"path" validate:"required"`

	// ReadOnly determines if the volume should only be read-only
	ReadOnly *bool `json:"read_only" validate:"required"`

	// Source is an incoming TaskVolume name that we want to use, in case we use multiple.
	Source string `json:"source" validate:"required"`
}

// TaskVolume defines the volumes that we want to add to the Task
type TaskVolume struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
