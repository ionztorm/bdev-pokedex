package pokecache

import (
	"sync"
	"time"
)

// create a Cache struct to hold
// 1) map[string]cacheEntry
// 2) a mutex to protect the map across goroutines.
// A cacheEntry should be a struct with two fields
// 1) createdAt - A time.Time that represents when the entry was created.
// 2) val - A []byte that represents the raw data we're caching.

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	cache    map[string]cacheEntry
	interval time.Duration
}

// Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Create a cache.Get() method that gets an entry from the cache.
// It should take a key (a string) and return a []byte and a bool.
// The bool should be true if the entry was found and false if it wasn't.

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record, exists := c.cache[key]
	if !exists {
		return nil, false
	}

	return record.val, true
}

// Create a cache.reapLoop() method that is called when the cache is created (by the NewCache function).
// Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval.
// This makes sure that the cache doesn't grow too large over time.
// For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.

func (c *Cache) reapLoop() {}

// expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).
func NewCache(interval time.Duration) *Cache {
	return &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}
}
