package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(defaultHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	return router
}
