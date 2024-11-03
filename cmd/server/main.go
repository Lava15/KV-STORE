package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/lava15/KV-STORE/internal/store"
	"github.com/lava15/KV-STORE/pkg/http"
	"log"
	net_http "net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	storeInstance := store.GetStore()
	handler := http.NewHandler(storeInstance)
	router := mux.NewRouter()
	router.Use(http.LoggingMiddleware)
	router.HandleFunc("/get", handler.GetHandler).Methods("GET")
	router.HandleFunc("/set", handler.SetHandler).Methods("POST")
	router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE")

	srv := &net_http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Starting server on port %s", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case err := <-serverErrors:
		log.Fatalf("Could not start server: %v", err)
	case sig := <-sigChan:
		log.Printf("Received signal %v. Shutting down...", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v", err)
		}
	}
}
