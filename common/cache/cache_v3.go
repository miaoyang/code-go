package cache

import (
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiration int64
}

type CacheV3 struct {
	items       sync.Map
	lock        sync.RWMutex
	defaultTTL  time.Duration
	maxCapacity int
	evictList   []interface{}
}

func NewCacheV3(defaultTTL time.Duration, maxCapacity int) *CacheV3 {
	return &CacheV3{
		defaultTTL:  defaultTTL,
		maxCapacity: maxCapacity,
		evictList:   make([]interface{}, 0, maxCapacity),
	}
}

func (c *CacheV3) Set(key string, value interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.cacheSize() >= c.maxCapacity {
		c.evict(1)
	}

	if ttl == 0 {
		ttl = c.defaultTTL
	}
	expiration := time.Now().Add(ttl).UnixNano()
	c.items.Store(key, &item{value, expiration})

	time.AfterFunc(ttl, func() {
		c.lock.Lock()
		defer c.lock.Unlock()

		if _, found := c.items.Load(key); found {
			c.items.Delete(key)
			c.evictList = append(c.evictList, key)
		}
	})
}

func (c *CacheV3) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if val, found := c.items.Load(key); found {
		item := val.(*item)
		if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
			c.items.Delete(key)
			return nil, false
		}
		return item.value, true
	}

	return nil, false
}

func (c *CacheV3) evict(count int) {
	for i := 0; i < count; i++ {
		key := c.evictList[0]
		c.evictList = c.evictList[1:]
		c.items.Delete(key)
	}
}

func (c *CacheV3) cacheSize() int {
	size := 0
	c.items.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
}
