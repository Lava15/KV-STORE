package main

import (
	"encoding/json"
	store "github.com/lava15/KV-STORE"
	"io"
	"log"
	"net/http"
)

var kvStore *store.Store

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	value, ok := kvStore.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	key, ok := req["key"]
	if !ok {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	value, ok := req["value"]
	if !ok {
		http.Error(w, "Missing value", http.StatusBadRequest)
		return
	}
	kvStore.Set(key, value)
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	kvStore.Delete(key)
	w.WriteHeader(http.StatusOK)
}

func main() {
	kvStore = store.NewStore()
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/delete", deleteHandler)

	log.Printf("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
