package client

import "errors"

var (
	ErrDockerTagFailed = errors.New("failed to validate dockerhub image")
)
