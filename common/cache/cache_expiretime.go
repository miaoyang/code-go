package cache

import (
	"sync"
	"time"
)

type CacheEx struct {
	data     map[string]cacheItem
	mu       sync.RWMutex
	expireCh chan string
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewCacheEx() *CacheEx {
	c := &CacheEx{
		data:     make(map[string]cacheItem),
		expireCh: make(chan string),
	}
	go c.startCleanup()
	return c
}

func (c *CacheEx) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expireTime := time.Now().Add(expiration)
	c.data[key] = cacheItem{
		value:      value,
		expiration: expireTime,
	}
	go c.scheduleExpiration(key, expireTime)
}

func (c *CacheEx) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if ok && item.expiration.After(time.Now()) {
		return item.value, true
	}
	return nil, false
}

func (c *CacheEx) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *CacheEx) startCleanup() {
	for {
		key := <-c.expireCh
		c.Delete(key)
	}
}

func (c *CacheEx) scheduleExpiration(key string, expireTime time.Time) {
	duration := time.Until(expireTime)
	timer := time.NewTimer(duration)
	<-timer.C
	c.expireCh <- key
}
