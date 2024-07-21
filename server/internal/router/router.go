package router

import (
	"lru-cache/internal/cache"
	"lru-cache/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(c *cache.LRUCache) *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/cache/{key}", handlers.GetHandler(c)).Methods("GET")
    r.HandleFunc("/cache", handlers.SetHandler(c)).Methods("POST")
    r.HandleFunc("/cache/{key}", handlers.DeleteHandler(c)).Methods("DELETE")
    r.HandleFunc("/ws", handlers.WebSocketHandler(c)).Methods("GET")
    return r
}
