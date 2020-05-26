package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

const FILE_PATH = ".files"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(defaultHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	return router
}
