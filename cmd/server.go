package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startServer() {
	router := NewRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", configuration.Port),
		Handler: router,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("starting server on port %s\n", configuration.Port)
	fmt.Println("intercepting endpoints:")
	for _, e := range configuration.Endpoints {
		fmt.Println("-", e)
	}

	<-done
	log.Println("server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		log.Println("good bye")
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Println("server exited properly")
}

func showBanner() {
	banner := `
███████ ██ ███    ███ ██████  ██      ███████     ██████  ███████ ███████ ████████     ███    ███  ██████   ██████ ██   ██
██      ██ ████  ████ ██   ██ ██      ██          ██   ██ ██      ██         ██        ████  ████ ██    ██ ██      ██  ██
███████ ██ ██ ████ ██ ██████  ██      █████       ██████  █████   ███████    ██        ██ ████ ██ ██    ██ ██      █████
     ██ ██ ██  ██  ██ ██      ██      ██          ██   ██ ██           ██    ██        ██  ██  ██ ██    ██ ██      ██  ██
███████ ██ ██      ██ ██      ███████ ███████     ██   ██ ███████ ███████    ██        ██      ██  ██████   ██████ ██   ██
                                                                                                           version %s
`
	fmt.Printf(banner, versionNumber)
}
