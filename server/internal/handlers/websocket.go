package handlers

import (
    "fmt"
    "net/http"
    "time"

    "lru-cache/internal/cache"
    "lru-cache/internal/structs"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func WebSocketHandler(c *cache.LRUCache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer conn.Close()

        for {
            c.Mutex.Lock()
            var currentItems []map[string]interface{}
            now := time.Now().UnixNano()

            for e := c.LRUList.Front(); e != nil; {
                next := e.Next()
                item := e.Value.(*structs.CacheItem)
                if item.Expiration > now {
                    currentItems = append(currentItems, map[string]interface{}{
                        "key":        item.Key,
                        "value":      item.Value,
                        "expiration": (item.Expiration - now) / int64(time.Second),
                    })
                } else {
                    // Remove expired item
                    c.LRUList.Remove(e)
                    delete(c.Cache, item.Key)
                }
                e = next
            }

            c.Mutex.Unlock()

            fmt.Println(currentItems)
            if err := conn.WriteJSON(currentItems); err != nil {
                return
            }

            time.Sleep(1 * time.Second)
        }
    }
}
