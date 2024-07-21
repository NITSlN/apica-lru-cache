package main

import (
    "net/http"
    "time"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func (c *LRUCache) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    for {
        c.mutex.Lock()
        currentItems := make(map[string]interface{})
        for key, elem := range c.cache {
            item := elem.Value.(*CacheItem)
            currentItems[key] = map[string]interface{}{
                "value":      item.value,
                "expiration": item.expiration,
            }
        }
        c.mutex.Unlock()

        if err := conn.WriteJSON(currentItems); err != nil {
            return
        }

        time.Sleep(1 * time.Second)
    }
}
