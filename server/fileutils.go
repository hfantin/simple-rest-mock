package server

import (
	"bytes"
	"crypto/tls"
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

const notFound = `{
	"httpCode": 404,
	"body":{}
}`

// type Response struct {
// 	HttpCode int                    `json:"httpCode"`
// 	Body     map[string]interface{} `json:"body,omitempty"`
// }

type Response struct {
	HttpCode int         `json:"httpCode"`
	Body     interface{} `json:"body,omitempty"`
}

const jsonContentType = "application/json"

func writeFileFromUrl(method, path string, headers http.Header, body []byte) {
	url := fmt.Sprintf("%s%s", config.Env.TargetServer, path)
	log.Printf("Writing %s of %s\n", method, url)
	var resp *http.Response
	var err error
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	switch method {
	case http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete:
		client := &http.Client{Transport: tr}
		req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
		req.Header = headers
		resp, err = client.Do(req)
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
	if len(body) > 0 {
		err := json.Unmarshal(body, &response.Body)
		if err != nil {
			log.Println("Failed to marshal body:", err)
			return
		}
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
	createFileIfNotExists(fileName)
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

func createFileIfNotExists(fileName string) {
	existe := fileExists(fileName)
	if !existe {
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			log.Printf("Failed to create new file %s:%s\n", fileName, err)
		}
		if _, err := file.Write([]byte(notFound)); err != nil {
			log.Printf("Failed to write data in the file %s: %s\n", fileName, err)
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getFileNameFromPath(path, method string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return fmt.Sprintf("%s/%s.%s.json", FILE_PATH, strings.ReplaceAll(path, "/", "."), method)
}
