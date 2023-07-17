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

type InsertResponse struct {
	Message string `json:"message"`
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

	h.store.Insert(words)

	res := InsertResponse{Message: "words are scheduled for insertion"}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *insertWordsHandler) validatePayload(p Payload) ([]string, error) {
	if !h.payloadRgx.MatchString(p.Words) {
		return nil, ErrBadPayloadFormat
	}

	words := strings.Split(p.Words, ",")

	for i := 0; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])
	}

	return words, nil
}

var ErrBadPayloadFormat = errors.New("words list must be of format: a,b,c")
