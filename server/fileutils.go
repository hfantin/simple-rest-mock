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

type Response struct {
	HttpCode int         `json:"httpCode"`
	Body     interface{} `json:"body,omitempty"`
}

func sendRequest(method, path string, headers http.Header, body []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", config.Env.TargetServer, path)
	log.Printf("sending %s on %s\n", method, url)
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
		err = errors.New("invalid or not implemented method: " + method)
	}
	if err != nil {
		return nil, err
	}
	return resp, err
}

func writeFile(path, method string, resp *http.Response) {
	body, _ := ioutil.ReadAll(resp.Body)
	filename := getFileNameFromPath(path, method)
	// for i, v := range resp.Header {
	// 	log.Printf("resp header %s with %s\n", i, v)
	// }
	var response Response
	contentType := resp.Header.Get("Content-type")
	// log.Printf("content-type %s\n", contentType)
	if len(body) > 0 && contentType == "application/json" {
		err := json.Unmarshal(body, &response.Body)
		if err != nil {
			log.Println("failed to marshal body:", err)
			return
		}
	} else {
		log.Printf("can't parse body: %v\n", &response.Body)
		return
	}
	response.HttpCode = resp.StatusCode
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("failed to create new file %s:%s\n", filename, err)
		return
	}
	defer file.Close()
	json, err := json.MarshalIndent(&response, "", "\t")
	if err != nil {
		log.Printf("failed to marshal json %s: %s\n", filename, err)
		return
	}
	if _, err := file.Write(json); err != nil {
		log.Printf("failed to write data in the file %s: %s\n", filename, err)
		return
	}
	// log.Printf("Data successfully recorded in the file %s %s %s\n!", filename, body, json)
	log.Printf("data successfully recorded in the file %s \n!", filename)

}

func readFile(path, method string) (*Response, error) {
	fileName := getFileNameFromPath(path, method)
	createFileIfNotExists(fileName)
	log.Printf("reading from %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var result *Response
	err = json.Unmarshal([]byte(b), &result)
	return result, err
}

func createFileIfNotExists(fileName string) {
	exists := fileExists(fileName)
	if !exists {
		file, err := os.Create(fileName)
		if err != nil {
			log.Printf("failed to create new file %s:%s\n", fileName, err)
		}
		defer file.Close()
		if _, err := file.Write([]byte(notFound)); err != nil {
			log.Printf("failed to write data in the file %s: %s\n", fileName, err)
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
