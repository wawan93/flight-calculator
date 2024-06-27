package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	"flights-test/config"
	"flights-test/internal/controllers/calculate"
	"flights-test/internal/services/calculator"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	calculatorService := &calculator.Calculator{}
	handler := calculate.NewCalculateController(calculatorService)

	root := mux.NewRouter()
	root.Handle("/calculate", handler)

	cfg, err := config.FromEnv()
	if err != nil {
		return fmt.Errorf("error parsing config: %w", err)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPListenPort),
		Handler: root,
	}

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(done)
	}()

	log.Println("starting server on port 8080")

	err = server.ListenAndServe()
	if err != http.ErrServerClosed {
		return fmt.Errorf("error starting server: %w", err)
	}

	<-done
	log.Println("server stopped")
	return nil
}
