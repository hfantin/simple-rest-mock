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
	path := r.URL.Path
	method := r.Method
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("Failed to read body")
	}
	if config.Env.WriteFile {
		writeFileFromUrl(method, path, r.Header, body)
	}
	log.Printf("Executing %s on %s with payload %s\n", method, path, body)
	response, err := readFile(path, method)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, map[string]interface{}{"error": fmt.Sprintf("%s", err)})
		return
	}
	if response.HttpCode >= http.StatusBadRequest {
		respondWithError(w, response.HttpCode, response.Body)
		return
	}
	respondWithJSON(w, response.HttpCode, response.Body)
}

func respondWithError(w http.ResponseWriter, code int, body map[string]interface{}) {
	respondWithJSON(w, code, body)
}

func respondWithJSON(w http.ResponseWriter, code int, body map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if len(body) > 0 {
		response, _ := json.Marshal(body)
		w.Write(response)
	}
}
