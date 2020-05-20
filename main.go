package main

import (
	"github.com/hfantin/simple-rest-mock/server"
)

func main() {
	// handleSigterm()
	server.StartWebServer()
}

// func handleSigterm() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	signal.Notify(c, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		log.Println("Stopping server...")
// 		os.Exit(1)
// 	}()
// }
