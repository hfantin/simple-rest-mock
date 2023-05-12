package cmd

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	url := fmt.Sprintf("%s%s", configuration.TargetServer, path)
	log.Printf("%s %s\n", method, url)
	var resp *http.Response
	var req *http.Request
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
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
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
	filename := getFileNameFromPath(path, method)
	var jsonResponse Response
	var body []byte
	contentType := resp.Header.Get("Content-type")
	contentEncoding := resp.Header.Get("Content-Encoding")
	defer resp.Body.Close()
	if contentEncoding == "gzip" {
		gReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			log.Println("can't uncompress body:", err)
			return
		}
		body, _ = io.ReadAll(gReader)
		gReader.Close()
	} else {
		body, _ = io.ReadAll(resp.Body)
	}
	if len(body) > 0 && strings.Contains(contentType, "application/json") {
		err := json.Unmarshal(body, &jsonResponse.Body)
		if err != nil {
			log.Println("failed to marshal body:", err)
			return
		}
	} else {
		log.Printf("can't parse body: %v\n", body)
		return
	}
	jsonResponse.HttpCode = resp.StatusCode
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("failed to create new file %s:%s\n", filename, err)
		return
	}
	defer file.Close()
	json, err := json.MarshalIndent(&jsonResponse, "", "\t")
	if err != nil {
		log.Printf("failed to marshal json %s: %s\n", filename, err)
		return
	}
	if _, err := file.Write(json); err != nil {
		log.Printf("failed to write data in the file %s: %s\n", filename, err)
		return
	}
	log.Printf("data successfully recorded in the file %s \n", filename)
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
	b, err := io.ReadAll(file)
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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonFolder := fmt.Sprintf("%s/.srm/%s", home, configuration.ResponseFilesPath)
	newPath := strings.TrimPrefix(path, "/")
	return fmt.Sprintf("%s/%s.%s.json", jsonFolder, strings.ReplaceAll(newPath, "/", "."), method)
}

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
