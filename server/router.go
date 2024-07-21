package main

import (
    "github.com/gorilla/mux"
)

func NewRouter(cache *LRUCache) *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/cache/{key}", cache.GetHandler).Methods("GET")
    r.HandleFunc("/cache", cache.SetHandler).Methods("POST")
    r.HandleFunc("/cache/{key}", cache.DeleteHandler).Methods("DELETE")
    r.HandleFunc("/ws", cache.WebSocketHandler)
    return r
}
