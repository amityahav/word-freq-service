package handlers

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

type healthCheckHandler struct{}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Status: "OK",
	}

	err := json.NewEncoder(w).Encode(&health)
	if err != nil {
		http.Error(w, "failed performing health check", http.StatusInternalServerError)
	}
}
