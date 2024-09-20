package middleware

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/knockbox/authentication/pkg/responses"
	"github.com/knockbox/matchbox/internal/client"
	"net/http"
)

var ActivityIdContextKey = "activity-id"

// ActivityId retrieves the activity from the database and stores the *models.Event with the ActivityIdContextKey
type ActivityId struct {
	l  hclog.Logger
	ec *client.EventClient
}

func (a *ActivityId) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activityId, ok := mux.Vars(r)["activity_id"]
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			responses.NewGenericError("activity_id was not provided").Encode(w)
			return
		}

		event, err := a.ec.GetByActivityId(activityId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			a.l.Error("failed to get activity by activity_id", "err", err)
			return
		}

		if event == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), ActivityIdContextKey, event)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UseActivityId(ec *client.EventClient, l hclog.Logger) *ActivityId {
	return &ActivityId{
		l:  l,
		ec: ec,
	}
}
