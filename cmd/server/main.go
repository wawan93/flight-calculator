package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	"flights-test/internal/controllers/calculate"
	"flights-test/internal/services/calculator"
)

func main() {
	calculatorService := &calculator.Calculator{}
	handler := calculate.NewCalculateController(calculatorService)

	root := mux.NewRouter()
	root.Handle("/calculate", handler)

	server := http.Server{
		Addr:    ":8080",
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
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-done

	log.Println("server stopped by user")
}
