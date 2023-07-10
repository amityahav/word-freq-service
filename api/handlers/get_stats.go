package handlers

import (
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

}
