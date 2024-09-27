package client

import "errors"

var (
	ErrDockerTagFailed        = errors.New("failed to validate dockerhub image")
	ErrDeploymentNotReady     = errors.New("the deployment is not ready")
	ErrDeploymentDoesNotExist = errors.New("the deployment does not exist")
	ErrVPCDoesNotExist        = errors.New("the deployment is missing a vpc")
	ErrClusterDoesNotExist    = errors.New("the deployment is missing a cluster")
	ErrTaskDefDoesNotExist    = errors.New("the deployment is missing a task definition")
)
