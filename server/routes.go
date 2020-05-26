package server

import (
	"github.com/gorilla/mux"
)

const FILE_PATH = ".files"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(defaultHandler)
	return router
}
