package pokecache

import (
	"testing"
	"time"
)

func TestCache_AddGetAndExpire(t *testing.T) {
	cacheInterval := 100 * time.Millisecond
	cache := NewCache(cacheInterval)

	key := "testkey"
	value := []byte("testvalue")

	cache.Add(key, value)

	got, exists := cache.Get(key)
	if !exists {
		t.Fatalf("expected key %q to exist in cache", key)
	}
	if string(got) != string(value) {
		t.Errorf("expected value %q, got %q", value, got)
	}

	// Wait for longer than cache interval to ensure entry expires
	time.Sleep(cacheInterval + 50*time.Millisecond)

	_, exists = cache.Get(key)
	if exists {
		t.Errorf("expected key %q to be expired and removed from cache", key)
	}
}
