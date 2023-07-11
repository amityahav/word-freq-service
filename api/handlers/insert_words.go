package handlers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
	"strings"
	"wordStore/internal"
)

type Payload struct {
	Words string `json:"words"`
}

type insertWordsHandler struct {
	store      *internal.Store
	payloadRgx *regexp.Regexp
}

func NewInsertWordsHandler(store *internal.Store) *insertWordsHandler {
	return &insertWordsHandler{
		store:      store,
		payloadRgx: regexp.MustCompile(`^(\w+)(,\s*\w+)*$`),
	}
}

func (h *insertWordsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p Payload

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	words, err := h.validatePayload(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.store.InsertWords(words)
}

func (h *insertWordsHandler) validatePayload(p Payload) ([]string, error) {
	if !h.payloadRgx.MatchString(p.Words) {
		return nil, errors.New("words list must be of format: a,b,c")
	}

	words := strings.Split(p.Words, ",")

	return words, nil
}
