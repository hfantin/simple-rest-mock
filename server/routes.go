package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

const FILE_PATH = ".files"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(defaultHandler)
	return router
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO get query parameters
	// TODO get body
	path := r.URL.Path
	method := r.Method
	log.Printf("Executing %s on %s", method, path)
	json, err := readFile(path)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("%s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, json)
}

// try to read file, if exists
func readFile(path string) (map[string]interface{}, error) {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	fileName := fmt.Sprintf("%s/%s.json", FILE_PATH, strings.ReplaceAll(path, "/", "."))
	log.Printf("trying to read %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)
	return result, err
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
