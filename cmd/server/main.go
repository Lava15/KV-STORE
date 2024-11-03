package main

import (
	"github.com/gorilla/mux"
	"github.com/lava15/KV-STORE/internal/store"
	"github.com/lava15/KV-STORE/pkg/http"
	"log"
	net_http "net/http"
)

func main() {
	storeInstance := store.GetStore()
	handler := http.NewHandler(storeInstance)
	router := mux.NewRouter()
	router.HandleFunc("/get", handler.GetHandler).Methods("GET")
	router.HandleFunc("/set", handler.SetHandler).Methods("POST")
	router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE")
	log.Printf("Listening on port 8080")
	log.Fatal(net_http.ListenAndServe(":8080", router))
}
