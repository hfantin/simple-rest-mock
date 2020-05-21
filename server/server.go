package server

import (
	"fmt"
	"net/http"
)

func New(port string) *http.Server {
	router := NewRouter()
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
}
