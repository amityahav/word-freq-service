package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"wordStore/api/handlers"
	"wordStore/internal"
)

const apiPrefix = "/api/v1"

// NewAPI creates a new http server for the service
func NewAPI(config *internal.Config) *http.Server {
	store := internal.NewStore(config.Store)

	// starting event-loop processing
	go store.Maintain()

	// init handlers
	hch := handlers.NewHealthCheckHandler()
	iwh := handlers.NewInsertWordsHandler(store)
	gsh := handlers.NewGetStatsHandler(store)

	// init routes
	router := mux.NewRouter().StrictSlash(true)

	router.Handle(apiPrefix+"/", hch).Methods("GET")
	router.Handle(apiPrefix+"/get_stats", gsh).Methods("GET")

	router.Handle(apiPrefix+"/insert_words", iwh).Methods("POST")

	server := http.Server{
		Addr:    config.ListenAddress,
		Handler: router,
	}

	return &server
}
