package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hfantin/simple-rest-mock/config"
)

const jsonContentType = "application/json"

func writeFileFromUrl(method, path string, body []byte) {
	url := fmt.Sprintf("%s%s", config.Env.TargetServer, path)
	log.Printf("Writing %s of %s with body %s\n", method, url, body)
	switch method {
	case "GET":
		resp, err := http.Get(url)
		writeFile(path, resp, err)
	case "POST":
		// ,
		resp, err := http.Post(url, jsonContentType, bytes.NewBuffer(body))
		writeFile(path, resp, err)
	// case "PUT":
	// 	resp, err := http.Put(url)
	// 	writeFile(resp, err)
	// case "DELETE":
	// 	resp, err := http.Delete(url)
	// 	writeFile(resp, err)
	default:
		log.Printf("Invalid method: %s\n", method)
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
		log.Printf("Failed to create new file %s:%s\n", filename, err)
		return
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		log.Printf("Failed to write data in the file %s: %s\n", filename, err)
		return
	}
	log.Printf("Data successfully recorded in the file %s\n!", filename)

}

// readFile function will try to read the file, if exists
func readFile(path string) (map[string]interface{}, error) {
	fileName := getFileNameFromPath(path)
	log.Printf("Reading from %s\n", fileName)
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
