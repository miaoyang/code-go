package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheNorm(t *testing.T) {
	cache := NewCache()

	// 设置缓存值
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// 读取缓存值
	value1, ok1 := cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 := cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)

	// 等待一段时间
	time.Sleep(5 * time.Second)

	// 再次读取缓存值
	value1, ok1 = cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 = cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)
}

func TestCacheExpireTime(t *testing.T) {
	cache := NewCacheEx()

	// 设置缓存值，带有过期时间
	cache.Set("key1", "value1", 2*time.Second)
	cache.Set("key2", "value2", 5*time.Second)

	// 读取缓存值
	value1, ok1 := cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 := cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)

	// 等待一段时间
	time.Sleep(3 * time.Second)

	// 再次读取缓存值
	value1, ok1 = cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 = cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)
}

func TestCacheLru(t *testing.T) {
	cache := NewLRUCache(5*time.Second, 2)

	// 设置缓存值，带有过期时间
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// 读取缓存值
	value1, ok1 := cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 := cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)

	// 等待一段时间
	time.Sleep(3 * time.Second)

	// 再次读取缓存值
	value1, ok1 = cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 = cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)
}

func TestCacheV3(t *testing.T) {
	c := NewCacheV3(time.Minute, 100)

	c.Set("key1", "value1", time.Second*30)
	c.Set("key2", "value2", time.Minute)

	val, found := c.Get("key1")
	if found {
		fmt.Println(val)
	}

	time.Sleep(time.Second * 45)

	val, found = c.Get("key1")
	if found {
		fmt.Println(val)
	}

	time.Sleep(time.Second * 30)

	val, found = c.Get("key1")
	if found {
		fmt.Println(val)
	} else {
		fmt.Println("key1 expired")
	}
}
