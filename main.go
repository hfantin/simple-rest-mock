package main

import (
	"github.com/hfantin/simple-rest-mock/config"
	"github.com/hfantin/simple-rest-mock/server"
)

func main() {
	// handleSigterm()
	server.New(config.Env.ServerPort)
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
