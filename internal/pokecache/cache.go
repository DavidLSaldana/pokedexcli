package pokecache

import (
	"sync"
	"time"
)

//in progress!

var time_interval int = 5

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() {
	interval := time.Duration(time_interval)

}

func (cache *Cache) Add(key string, val []byte) {
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, ok
}
