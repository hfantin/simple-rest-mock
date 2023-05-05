package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hfantin/simple-rest-mock/config"
	"github.com/hfantin/simple-rest-mock/server"
)

func main() {
	srv := server.New()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go startServer(srv)
	log.Printf("Server Started on port %s - https enabled? %t\n", config.Env.ServerPort, config.Env.UseHTTPS)

	<-done
	log.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here if necessary
		log.Println("Good bye")
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server Exited Properly")
}

func startServer(srv *http.Server) {
	var err error
	if config.Env.UseHTTPS {
		err = srv.ListenAndServeTLS(config.Env.CertificatePath, config.Env.KeyPath)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
