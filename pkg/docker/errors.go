package docker

import "errors"

var (
	ErrRateLimitExceeded    = errors.New("rate-limit exceeded try again later")
	ErrUnexpectedStatusCode = errors.New("an unexpected status code was returned by hub.docker.com")
)
