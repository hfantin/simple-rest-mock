package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hfantin/simple-rest-mock/config"
)

type Response struct {
	HttpCode int                    `json:"httpCode"`
	Body     map[string]interface{} `json:"body"`
}

const jsonContentType = "application/json"

func writeFileFromUrl(method, path string, body []byte) {
	url := fmt.Sprintf("%s%s", config.Env.TargetServer, path)
	log.Printf("Writing %s of %s with body %s\n", method, url, body)
	var resp *http.Response
	var err error
	switch method {
	case http.MethodGet:
		resp, err = http.Get(url)
	case http.MethodPost:
		resp, err = http.Post(url, jsonContentType, bytes.NewBuffer(body))
	default:
		err = errors.New("Invalid or not implemented method: " + method)
	}
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return
	}
	writeFile(path, method, resp)
}

func writeFile(path, method string, resp *http.Response) {

	body, _ := ioutil.ReadAll(resp.Body)
	filename := getFileNameFromPath(path, method)
	var response Response
	err := json.Unmarshal(body, &response.Body)
	if err != nil {
		log.Println("Failed to marshal body:", err)
		return
	}
	response.HttpCode = resp.StatusCode
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Printf("Failed to create new file %s:%s\n", filename, err)
		return
	}
	json, err := json.MarshalIndent(&response, "", "\t")
	if err != nil {
		log.Printf("Failed to marshal json %s: %s\n", filename, err)
		return
	}
	if _, err := file.Write(json); err != nil {
		log.Printf("Failed to write data in the file %s: %s\n", filename, err)
		return
	}
	log.Printf("Data successfully recorded in the file %s %s %s\n!", filename, body, json)

}

func readFile(path, method string) (*Response, error) {
	fileName := getFileNameFromPath(path, method)
	log.Printf("Reading from %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	var result *Response
	err = json.Unmarshal([]byte(b), &result)
	return result, err
}

func getFileNameFromPath(path, method string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return fmt.Sprintf("%s/%s.%s.json", FILE_PATH, strings.ReplaceAll(path, "/", "."), method)
}
