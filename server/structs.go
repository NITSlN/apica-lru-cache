package main

import (
	"container/list"
	"sync"
)


type CacheItem struct {
    key        string
    value      interface{}
    expiration int64
}

type LRUCache struct {
    cache   map[string]*list.Element
    lruList *list.List
    mutex   sync.Mutex
}