package caching

import (
	"sync"
	"time"
)

type Cache struct {
	cacheItem map[string]cacheEntry
	mux       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheItem: make(map[string]cacheEntry),
		mux:       &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c

}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cacheItem[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cacheItem[key]
	if !ok {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, entry := range c.cacheItem {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.cacheItem, key)
		}
	}
}
