// Package pokecache provides a simple in-memory caching mechanism
// with automatic expiration of entries based on a configured interval.
package pokecache

import (
	"sync"
	"time"
)

type Cacher interface {
	Get(key string) ([]byte, bool)
	Add(key string, val []byte)
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	cache    map[string]cacheEntry
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record, exists := c.cache[key]
	if !exists {
		return nil, false
	}

	return record.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()

		for key, val := range c.cache {
			if time.Since(val.createdAt) > c.interval {
				delete(c.cache, key)

			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}
