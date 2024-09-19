package docker

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

const ENDPOINT = "https://hub.docker.com/v2"

type Client struct {
	*http.Client
	limit *rate.Limiter
	l     hclog.Logger
}

// CheckRepositoryTag checks to see if a tag exists for a given repository of the provided namespace.
// see: https://docs.docker.com/reference/api/hub/latest/#tag/repositories/paths/~1v2~1namespaces~1%7Bnamespace%7D~1repositories~1%7Brepository%7D~1tags~1%7Btag%7D/head
func (c *Client) CheckRepositoryTag(ctx context.Context, options *CheckRepositoryTagOptions) *CheckRepositoryTagResult {
	result := &CheckRepositoryTagResult{}
	url := fmt.Sprintf("%s/namespaces/%s/repositories/%s/tags/%s", ENDPOINT, options.Namespace, options.Repository, options.Tag)

	// Rate-Limit handling.
	if c.limit != nil {
		if err := c.limit.Wait(ctx); err != nil {
			result.Error = err
			return result
		}
	}

	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		c.l.Error("CheckRepositoryTag failed to create request", "error", err)
		result.Error = err
		return result
	}

	res, err := c.Do(req)
	if err != nil {
		c.l.Info("CheckRepositoryTag responded with an error", "error", err)
		result.Error = err
		return result
	}

	// If we haven't set up our limiter, we need ot do so.
	if c.limit == nil {
		rLimit := res.Header.Get("X-RateLimit-Limit")
		limit, err := strconv.Atoi(rLimit)
		if err != nil {
			c.l.Error("CheckRepositoryTag failed to parse X-RateLimit-Limit")
			result.Error = err
			return result
		}

		c.limit = rate.NewLimiter(rate.Every(1*time.Minute), limit)

		rRemaining := res.Header.Get("X-RateLimit-Remaining")
		remaining, err := strconv.Atoi(rRemaining)
		if err != nil {
			c.l.Error("CheckRepositoryTag failed to parse X-RateLimit-Remaining")
			result.Error = err
			return result
		}

		// Take away any tokens we don't currently have.
		if remaining != limit {
			_ = c.limit.ReserveN(time.Now().Add(1*time.Minute), limit-remaining)
		}

		c.l.Info("docker.Client Limiter created", "limit", limit, "remaining", remaining)
	}

	switch res.StatusCode {
	case 200:
		result.Exists = true
		return result
	case 403:
		result.Exists = true
		result.Private = true
		return result
	case 404:
		return result
	case 429:
		// Consume all remaining tokens if we are currently rate-limited.
		if c.limit != nil {
			_ = c.limit.ReserveN(time.Now().Add(1*time.Minute), int(c.limit.Tokens()))
		}
		result.Error = ErrRateLimitExceeded
		return result
	default:
		result.Error = ErrUnexpectedStatusCode
		return result
	}
}

func NewClient(l hclog.Logger) *Client {
	return &Client{
		Client: http.DefaultClient,
		limit:  nil,
		l:      l,
	}
}
