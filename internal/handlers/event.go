package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/knockbox/authentication/pkg/enums"
	"github.com/knockbox/authentication/pkg/middleware"
	"github.com/knockbox/authentication/pkg/responses"
	utils2 "github.com/knockbox/authentication/pkg/utils"
	"github.com/knockbox/matchbox/internal/client"
	"github.com/knockbox/matchbox/pkg/docker"
	"github.com/knockbox/matchbox/pkg/models"
	"github.com/knockbox/matchbox/pkg/payloads"
	"github.com/knockbox/matchbox/pkg/utils"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
)

type Event struct {
	l  hclog.Logger
	ec *client.EventClient
}

func (e *Event) Create(w http.ResponseWriter, r *http.Request) {
	payload := &payloads.EventCreate{}
	if utils2.DecodeAndValidateStruct(w, r, payload) {
		return
	}

	token := *r.Context().Value(middleware.BearerTokenContextKey).(*jwt.Token)
	accountId, _, role := utils.ParseUserClaims(token)

	// Must be a User to create an Event.
	if !role.HasRequiredRole(enums.User) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := e.ec.CreateEvent(payload, accountId); err != nil {
		if utils2.IsDuplicateEntry(err) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			msg := "an event with the provided name already exists"
			responses.NewGenericError(msg).Encode(w)
			return
		} else if errors.Is(err, docker.ErrUnexpectedStatusCode) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			responses.NewGenericError(err.Error()).Encode(w)
			return
		} else if errors.Is(err, docker.ErrRateLimitExceeded) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			responses.NewGenericError(err.Error()).Encode(w)
			return
		} else if errors.Is(err, client.ErrDockerTagFailed) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			responses.NewGenericError(err.Error()).Encode(w)
			return
		}

		http.Error(w, "failed to create event", http.StatusInternalServerError)
		e.l.Error("failed to create event", "error", err, "payload", payload)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (e *Event) GetAll(w http.ResponseWriter, r *http.Request) {
	events, err := e.ec.GetAllEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e.l.Error("failed to get events", "err", err)
		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var dtos []*models.EventDTO
	for _, event := range events {
		dtos = append(dtos, event.DTO())
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dtos)
}

func (e *Event) GetByActivityId(w http.ResponseWriter, r *http.Request) {
	activityId, ok := mux.Vars(r)["activity_id"]
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		responses.NewGenericError("activity_id was not provided").Encode(w)
		return
	}

	event, err := e.ec.GetByActivityId(activityId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e.l.Error("failed to get user by account_id", "err", err)
		return
	}

	if event == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(event.DTO())
}

func (e *Event) Route(r *mux.Router) {
	eventRouter := r.PathPrefix("/events").Subrouter()
	eventRouter.HandleFunc("", e.Create).Methods(http.MethodPost)
	eventRouter.HandleFunc("", e.GetAll).Methods(http.MethodGet)
	eventRouter.HandleFunc("/{activity_id}", e.GetByActivityId).Methods(http.MethodGet)
}

func NewEvent(l hclog.Logger) *Event {
	db, err := utils2.MySQLConnection()
	if err != nil {
		panic(err)
	}

	return &Event{
		l:  l,
		ec: client.NewEventClient(db, l),
	}
}
