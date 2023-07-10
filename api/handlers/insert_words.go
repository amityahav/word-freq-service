package handlers

import (
	"net/http"
	"wordStore/internal"
)

type insertWordsHandler struct {
	store *internal.Store
}

func NewInsertWordsHandler(store *internal.Store) *insertWordsHandler {
	return &insertWordsHandler{store: store}
}

func (h *insertWordsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
