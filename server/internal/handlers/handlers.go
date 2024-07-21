package handlers

import (
	"encoding/json"
	"lru-cache/internal/cache"
	"lru-cache/internal/structs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        key := vars["key"]

        value, found := c.Get(key)
        if !found {
            http.Error(w, "Key not found", http.StatusNotFound)
            return
        }

        json.NewEncoder(w).Encode(value)
    }
}

func SetHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var requestData structs.SetRequest

        if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        c.Set(requestData.Key, requestData.Value, time.Duration(requestData.Expiration)*time.Second)
        w.WriteHeader(http.StatusOK)
    }
}

func DeleteHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        key := vars["key"]

        c.Delete(key)
        w.WriteHeader(http.StatusOK)
    }
}
