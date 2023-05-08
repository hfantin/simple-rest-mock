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

// TODO dinamic list
// var endpoints []string = []string{"/resource/calculoValorBloquear/calcular"}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	endpoint := r.URL.Path
	method := r.Method

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("failed to read body")
	}
	isEndpointIntercepted := contains(config.Env.Endpoints, endpoint)
	if isEndpointIntercepted {
		log.Printf("intercepting %s request\n", endpoint)
		if config.Env.WriteFile {
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
			respondWithError(w, resp.StatusCode, resp.Body)
		}
		for k := range resp.Header {
			w.Header().Set(k, resp.Header.Get(k))
		}
		w.WriteHeader(resp.StatusCode)
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
