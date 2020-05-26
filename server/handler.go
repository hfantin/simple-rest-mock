package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hfantin/simple-rest-mock/config"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO get query parameters
	// TODO get body
	// try to read from target
	path := r.URL.Path
	method := r.Method
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("Failed to read body")
	}
	if config.Env.WriteFile {
		writeFileFromUrl(method, path, body)
	}
	log.Printf("Executing %s on %s\n", method, path)
	response, err := readFile(path, method)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	if response.HttpCode >= http.StatusBadRequest {
		// TODO convert error map to string
		log.Printf("Error from file: %s", response.Body)
		respondWithError(w, response.HttpCode, "Error")
		return
	}

	respondWithJSON(w, response.HttpCode, response.Body)
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
