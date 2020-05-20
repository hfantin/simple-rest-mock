package server

import (
	"fmt"
	"log"
	"net/http"
)

func New(port string) {
	router := Router()
	log.Println("Serving on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
