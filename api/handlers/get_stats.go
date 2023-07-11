package handlers

import (
	"encoding/json"
	"net/http"
	"wordStore/internal"
)

type getStatsHandler struct {
	store *internal.Store
}

func NewGetStatsHandler(store *internal.Store) *getStatsHandler {
	return &getStatsHandler{store: store}
}

func (h *getStatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	stats, err := h.store.GetStats(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
