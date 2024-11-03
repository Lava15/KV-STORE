package http

import (
	"encoding/json"
	"github.com/lava15/KV-STORE/internal/store"
	"io"
	"net/http"
)

type Handler struct {
	Store *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "missing key", http.StatusBadRequest)
		return
	}
	value, ok := h.Store.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (h *Handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid requst bodu", http.StatusBadRequest)
		return
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
	h.Store.Set(key, value)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "missing key", http.StatusBadRequest)
		return
	}
	h.Store.Delete(key)
	w.WriteHeader(http.StatusOK)
}
