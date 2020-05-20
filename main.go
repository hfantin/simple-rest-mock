package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hfantin/simple-rest-mock/server"
)

func main() {
	handleSigterm()
	go server.StartWebServer()
	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()

}

func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Stopping server...")
		os.Exit(1)
	}()
}
