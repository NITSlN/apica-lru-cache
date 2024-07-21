package main

import (
	"log"
	"lru-cache/internal/cache"
	"lru-cache/internal/router"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

func main() {
	cache := cache.NewLRUCache()
	r := router.NewRouter(cache)

	allowedOrigins := []string{
        "http://localhost:3000",
    }
	corsHandler := handlers.CORS(
        handlers.AllowedOrigins(allowedOrigins),
        handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )(r)

	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
