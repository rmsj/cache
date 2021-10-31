package cache

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type cacheEntry struct {
	key       int
	value     int
	timestamp time.Time
}

type LRUCache struct {
	data []cacheEntry
	mu   sync.Mutex
}

// NewLRUCache builds and returns a pointer to an LRUCache
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity < 2 {
		return nil, errors.New("capacity must be at least 2")
	}
	c := make([]cacheEntry, 0, capacity)
	lru := LRUCache{
		data: c,
	}

	return &lru, nil
}

//Put adds a value to the cache with specified key
func (c *LRUCache) Put(key int, value int) error {

	c.mu.Lock()
	{
		// checks if we reached capacity
		if len(c.data) == cap(c.data) {
			// we know the last one is the least used
			c.data = c.data[:len(c.data)-1]
		}

		// new cache entry
		c.data = append(c.data, cacheEntry{
			key:       key,
			value:     value,
			timestamp: time.Now(),
		})

		// send least used/oldest to end
		sort.SliceStable(c.data, func(i, j int) bool {
			return c.data[i].timestamp.After(c.data[j].timestamp)
		})
	}

	c.mu.Unlock()

	return nil
}

//Get gets a value out of the cache or -1 if not found
func (c *LRUCache) Get(key int) int {

	idx := -1
	value := -1
	c.mu.Lock()
	{
		for i, v := range c.data {
			if v.key == key {
				idx = i
				break
			}
		}

		if idx != -1 {
			c.data[idx].timestamp = time.Now()
			value = c.data[idx].value
		}

		// reorder
		sort.SliceStable(c.data, func(i, j int) bool {
			return c.data[i].timestamp.After(c.data[j].timestamp)
		})
	}

	c.mu.Unlock()

	return value
}
