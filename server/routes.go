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
	"github.com/hfantin/simple-rest-mock/config"
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
	// try to read from target
	path := r.URL.Path
	method := r.Method
	if config.Env.WriteFile {
		writeFileFromUrl(method, path)
	}
	log.Printf("Executing %s on %s", method, path)
	json, err := readFile(path)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, json)
}

func writeFileFromUrl(method, path string) {
	url := fmt.Sprintf("%s%s", config.Env.TargetServer, path)
	log.Printf("writing %s of %s\n", method, url)
	switch method {
	case "GET":
		resp, err := http.Get(url)
		writeFile(path, resp, err)
	// case "POST":
	// 	resp, err := http.Post(url)
	// 	writeFile(resp, err)
	// case "PUT":
	// 	resp, err := http.Put(url)
	// 	writeFile(resp, err)
	// case "DELETE":
	// 	resp, err := http.Delete(url)
	// 	writeFile(resp, err)
	default:
		log.Printf("Invalid method: %s", method)
	}
}

func writeFile(path string, resp *http.Response, err error) {
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	filename := getFileNameFromPath(path)
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Failed to create new file %s:%s", filename, err)
		return
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		log.Printf("Failed to write data in the file %s: %s", filename, err)
		return
	}
	log.Printf("Data successfully recorded in the file %s\n!", filename)

}

// readFile function will try to read the file, if exists
func readFile(path string) (map[string]interface{}, error) {
	fileName := getFileNameFromPath(path)
	log.Printf("trying to read %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(b), &result)
	return result, err
}

func getFileNameFromPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return fmt.Sprintf("%s/%s.json", FILE_PATH, strings.ReplaceAll(path, "/", "."))
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
