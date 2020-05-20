package server

import (
	"log"
	"net/http"
)

func StartWebServer() {
	router := Router()
	log.Println("Serving on port", 5000)
	err := http.ListenAndServe(":5000", router)
	// err := http.ListenAndServe(":5000"+utils.Env.ServerPort, router)
	if err != nil {
		// log.Println("An error occured starting HTTP listener at port", utils.Env.ServerPort)
		log.Println("Error: " + err.Error())
	}
}
