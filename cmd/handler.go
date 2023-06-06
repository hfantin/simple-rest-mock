package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	endpoint := r.URL.Path
	method := r.Method

	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("failed to read body")
	}
	isEndpointIntercepted := contains(configuration.Endpoints, endpoint)
	if isEndpointIntercepted {
		log.Printf("intercepting %s request\n", endpoint)
		if configuration.RecMode {
			resp, err := sendRequest(method, path, r.Header, body)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
				return
			}
			writeFile(endpoint, method, resp)
		}
		response, err := readFile(endpoint, method)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, map[string]interface{}{"error": fmt.Sprintf("%s", err)})
			return
		}
		if response.HttpCode >= http.StatusBadRequest {
			respondWithError(w, response.HttpCode, response.Body)
			return
		}
		respondWithJSON(w, response.HttpCode, response.Body)
	} else {
		resp, err := sendRequest(method, path, r.Header, body)
		if err != nil {
			log.Printf("failed to send request %s:%v\n", path, err)
			respondWithError(w, resp.StatusCode, resp.Body)
			return
		}
		for k := range resp.Header {
			w.Header().Set(k, resp.Header.Get(k))
		}
		io.Copy(w, resp.Body)
	}

}

func respondWithError(w http.ResponseWriter, code int, body interface{}) {
	respondWithJSON(w, code, body)
}

func respondWithJSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response, _ := json.Marshal(body)
	w.Write(response)
}
