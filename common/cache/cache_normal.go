package cache

import (
	"sync"
)

type CacheNormal struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *CacheNormal {
	return &CacheNormal{
		data: make(map[string]interface{}),
	}
}

func (c *CacheNormal) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *CacheNormal) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}
