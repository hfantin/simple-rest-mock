package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(DefaultHandler)
	return router
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO get query parameters
	// TODO get body
	path := r.URL.Path
	method := r.Method
	log.Printf("Executing %s on %s", method, path)
	// TODO based on path, read and return the file content if exists
	RespondWithJSON(w, http.StatusOK, "OK")
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
