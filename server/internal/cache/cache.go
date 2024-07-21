package cache

import (
	"container/list"
	"lru-cache/internal/structs"
	"sync"
	"time"
)

type LRUCache struct {
    Cache   map[string]*list.Element
    LRUList *list.List
    Mutex   sync.Mutex
}

func NewLRUCache() *LRUCache {
    return &LRUCache{
        Cache:   make(map[string]*list.Element),
        LRUList: list.New(),
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {

    c.Mutex.Lock()
    defer c.Mutex.Unlock()

    if elem, found := c.Cache[key]; found {
        item := elem.Value.(*structs.CacheItem)
        if time.Now().UnixNano() > item.Expiration {
            c.LRUList.Remove(elem)
            delete(c.Cache, key)
            return nil, false
        }
        c.LRUList.MoveToFront(elem)
        return item.Value, true
    }
    return nil, false
}

func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
    c.Mutex.Lock()
    defer c.Mutex.Unlock()

    if elem, found := c.Cache[key]; found {
        c.LRUList.MoveToFront(elem)
        elem.Value.(*structs.CacheItem).Value = value
        elem.Value.(*structs.CacheItem).Expiration = time.Now().Add(duration).UnixNano()
        return
    }

    item := &structs.CacheItem{
        Key:        key,
        Value:      value,
        Expiration: time.Now().Add(duration).UnixNano(),
    }
    elem := c.LRUList.PushFront(item)
    c.Cache[key] = elem
}

func (c *LRUCache) Delete(key string) {
    c.Mutex.Lock()
    defer c.Mutex.Unlock()

    if elem, found := c.Cache[key]; found {
        c.LRUList.Remove(elem)
        delete(c.Cache, key)
    }
}
