package handlers

import (
	"lru-cache/internal/cache"
	"lru-cache/internal/structs"
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
            currentItems := make(map[string]interface{})
            now := time.Now().UnixNano()
            for key, elem := range c.Cache {
                item := elem.Value.(*structs.CacheItem)
				if item.Expiration > now {
					currentItems[key] = map[string]interface{}{
						"value":      item.Value,
						"expiration": (item.Expiration - now) / int64(time.Second),
					}
				} else {
					// Remove expired item
					c.LRUList.Remove(elem)
					delete(c.Cache, key)
				}
            }
            c.Mutex.Unlock()

            if err := conn.WriteJSON(currentItems); err != nil {
                return
            }

            time.Sleep(1 * time.Second)
        }
    }
}