package server

import (
	"fmt"
	"net/http"

	"github.com/hfantin/simple-rest-mock/config"
)

func New() *http.Server {
	router := NewRouter()
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.ServerPort),
		Handler: router,
	}
}
