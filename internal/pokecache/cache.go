package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache

}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.entries[key]
	return entry.val, ok
}

func (cache *Cache) reap(currentTime time.Time, last time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, value := range cache.entries {
		if value.createdAt.Before(currentTime.Add(-last)) {
			delete(cache.entries, key)
		}
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}

}
