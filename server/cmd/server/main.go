package main

import (
	"log"
	"lru-cache/internal/cache"
	"lru-cache/internal/router"
	"net/http"
	"time"
)

func main() {
	cache := cache.NewLRUCache()
	r := router.NewRouter(cache)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
