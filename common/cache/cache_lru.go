package cache

import (
	"container/list"
	"sync"
	"time"
)

const LRUCACHE_MAX_LEN = 2

type CacheItem struct {
	key        string
	value      interface{}
	expiration time.Time
}

type LRUCache struct {
	cache      map[string]*list.Element
	evictList  *list.List
	mu         sync.RWMutex
	expiration time.Duration
}

func NewLRUCache(expiration time.Duration, maxSize int) *LRUCache {
	return &LRUCache{
		cache:      make(map[string]*list.Element),
		evictList:  list.New(),
		expiration: expiration,
	}
}

func (c *LRUCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.evictList.MoveToFront(elem)
		elem.Value.(*CacheItem).value = value
		elem.Value.(*CacheItem).expiration = time.Now().Add(c.expiration)
	} else {
		expiration := time.Now().Add(c.expiration)
		item := &CacheItem{key, value, expiration}
		elem := c.evictList.PushFront(item)
		c.cache[key] = elem
		if c.evictList.Len() > LRUCACHE_MAX_LEN {
			c.removeOldest()
		}
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		item := elem.Value.(*CacheItem)
		if item.expiration.After(time.Now()) {
			c.evictList.MoveToFront(elem)
			return item.value, true
		} else {
			c.removeElement(elem)
		}
	}
	return nil, false
}

func (c *LRUCache) removeOldest() {
	elem := c.evictList.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}

func (c *LRUCache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	item := e.Value.(*CacheItem)
	delete(c.cache, item.key)
}
