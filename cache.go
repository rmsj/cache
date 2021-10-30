package cache

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type cacheEntry struct{
	key interface{}
	value interface{}
	timestamp time.Time
}

type LRUCache struct{
	data []cacheEntry
	mu sync.Mutex
}

// NewLRUCache builds and returns a pointer to an LRUCache
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity < 2 {
		return nil, errors.New("capacity must be at least 2")
	}
	c := make([]cacheEntry, capacity, capacity)
	lru := LRUCache{
		data: c,
	}

	return &lru, nil
}

func (c *LRUCache) put(key interface{}, value interface{}) error {

	c.mu.Lock()
	{
		// checks if we reached capacity
		if len(c.data) == cap(c.data) {
			// we know the last one is the least used
			c.data = c.data[:len(c.data)-1]
		}

		// new cache entry
		entry := cacheEntry{
			key: key,
			value: value,
			timestamp: time.Now(),
		}

		c.data = append(c.data, entry)

		sort.Slice()
	}

	c.mu.Unlock()

	return nil
}