package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Healthcheck struct{}

func (h *Healthcheck) GetHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]bool{"OK": true})
}

func (h *Healthcheck) Route(r *mux.Router) {
	r.HandleFunc("/health", h.GetHealthcheck).Methods(http.MethodGet)
}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}
