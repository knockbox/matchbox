package handlers

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/knockbox/authentication/pkg/responses"
	"github.com/knockbox/matchbox/pkg/docker"
	"net/http"
)

type Docker struct {
	c *docker.Client
	l hclog.Logger
}

func (d *Docker) IsValidRepository(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := d.c.CheckRepositoryTag(r.Context(), &docker.CheckRepositoryTagOptions{
		Namespace:  vars["namespace"],
		Repository: vars["repository"],
		Tag:        vars["tag"],
	})
	if result.Error != nil {
		if errors.Is(result.Error, docker.ErrUnexpectedStatusCode) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			responses.NewGenericError(result.Error.Error()).Encode(w)
			return
		} else if errors.Is(result.Error, docker.ErrRateLimitExceeded) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			responses.NewGenericError(result.Error.Error()).Encode(w)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If the repository is private, we need the user to make it public.
	if result.Private {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		responses.NewGenericError("the resource is currently private, please make it public and try again").Encode(w)
		return
	}

	// If the repository doesn't exist, we need the user to fix their inputs.
	if !result.Exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// All is good, we can continue.
	w.WriteHeader(http.StatusNoContent)
}

func (d *Docker) Route(r *mux.Router) {
	dockerRouter := r.PathPrefix("/docker").Subrouter()
	dockerRouter.HandleFunc("/{namespace}/{repository}/{tag}", d.IsValidRepository).Methods(http.MethodGet)
}

func NewDocker(l hclog.Logger) *Docker {
	return &Docker{
		c: docker.NewClient(l),
		l: l,
	}
}
