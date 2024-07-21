package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "log"
    "github.com/gorilla/mux"
    "lru-cache/internal/cache"
    "lru-cache/internal/structs"
)


// GetHandler handles the GET requests to retrieve a value from the cache.
func GetHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()
        vars := mux.Vars(r)
        key := vars["key"]

        value, found := c.Get(key)
        w.Header().Set("Content-Type", "application/json")
        if !found {
            log.Printf("Key not found: %s", key)
            resp := structs.Response{
                Status:  "error",
                Message: "Key not found",
            }
            w.WriteHeader(http.StatusNotFound)
            json.NewEncoder(w).Encode(resp)
            return
        }

        resp := structs.Response{
            Status:  "success",
            Message: "Key found",
            Data:    value,
        }
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(resp); err != nil {
            log.Printf("Error encoding response: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
    }
}

// SetHandler handles the POST requests to set a value in the cache.
func SetHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()

        var requestData structs.SetRequest

        if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
            log.Printf("Error decoding request: %v", err)
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        c.Set(requestData.Key, requestData.Value, time.Duration(requestData.Expiration)*time.Second)
        resp := structs.Response{
            Status:  "success",
            Message: "Key set successfully",
        }
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(resp); err != nil {
            log.Printf("Error encoding response: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
    }
}

// DeleteHandler handles the DELETE requests to remove a value from the cache.
func DeleteHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()
        vars := mux.Vars(r)
        key := vars["key"]

        c.Delete(key)
        resp := structs.Response{
            Status:  "success",
            Message: "Key deleted successfully",
        }
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(resp); err != nil {
            log.Printf("Error encoding response: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
    }
}
