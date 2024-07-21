package main

import (
    "container/list"
    "time"
)


func NewLRUCache() *LRUCache {
    return &LRUCache{
        cache:   make(map[string]*list.Element),
        lruList: list.New(),
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, found := c.cache[key]; found {
        item := elem.Value.(*CacheItem)
        if time.Now().UnixNano() > item.expiration {
            c.lruList.Remove(elem)
            delete(c.cache, key)
            return nil, false
        }
        c.lruList.MoveToFront(elem)
        return item.value, true
    }
    return nil, false
}

func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, found := c.cache[key]; found {
        c.lruList.MoveToFront(elem)
        elem.Value.(*CacheItem).value = value
        elem.Value.(*CacheItem).expiration = time.Now().Add(duration).UnixNano()
        return
    }

    item := &CacheItem{
        key:        key,
        value:      value,
        expiration: time.Now().Add(duration).UnixNano(),
    }
    elem := c.lruList.PushFront(item)
    c.cache[key] = elem
}

func (c *LRUCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, found := c.cache[key]; found {
        c.lruList.Remove(elem)
        delete(c.cache, key)
    }
}
