package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

func (c *LRUCache) GetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["key"]

    value, found := c.Get(key)
    if !found {
        http.Error(w, "Key not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(value)
}

func (c *LRUCache) SetHandler(w http.ResponseWriter, r *http.Request) {
    var requestData struct {
        Key        string      `json:"key"`
        Value      interface{} `json:"value"`
        Expiration int         `json:"expiration"` // in seconds
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    c.Set(requestData.Key, requestData.Value, time.Duration(requestData.Expiration)*time.Second)
    w.WriteHeader(http.StatusOK)
}

func (c *LRUCache) DeleteHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["key"]

    c.Delete(key)
    w.WriteHeader(http.StatusOK)
}
