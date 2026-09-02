package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu         *sync.RWMutex
	cached     map[string]cacheEntry
	defaultTTL time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func AppCache() *Cache {
	c := &Cache{
		mu:         &sync.RWMutex{},
		cached:     make(map[string]cacheEntry),
		defaultTTL: 5 * time.Minute,
	}

	go c.reapLoop(5 * time.Minute)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cached[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	cached, found := c.cached[key]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}

	c.mu.RUnlock()
	return cached.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		now := time.Now()
		for key, val := range c.cached {
			if now.Sub(val.createdAt) > c.defaultTTL {
				delete(c.cached, key)
			}
		}
		c.mu.Unlock()
	}
}
