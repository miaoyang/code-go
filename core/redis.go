package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// https://zhuanlan.zhihu.com/p/637537337
// https://blog.csdn.net/qq_44237719/article/details/128920821

var Redis *RedisClient

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

// InitRedis 初始化redis
func InitRedis() {
	redisConfig := Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	redisClient := &RedisClient{
		client: client,
		ctx:    context.Background(),
	}
	Redis = redisClient
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		LOG.Println("redis连接失败！")
		return
	}
	LOG.Println("redis连接成功！")
}

/*------------------------------------ 字符 操作 ------------------------------------*/

// Set 设置 key的值
func (this *RedisClient) Set(key, value string) bool {
	result, err := this.client.Set(this.ctx, key, value, 0).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return result == "OK"
}

// SetEX 设置 key的值并指定过期时间
func (this *RedisClient) SetEX(key, value string, ex time.Duration) bool {
	result, err := this.client.Set(this.ctx, key, value, ex).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return result == "OK"
}

// Get 获取 key的值
func (this *RedisClient) Get(key string) (bool, string) {
	result, err := this.client.Get(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
		return false, ""
	}
	return true, result
}

// GetSet 设置新值获取旧值
func (this *RedisClient) GetSet(key, value string) (bool, string) {
	oldValue, err := this.client.GetSet(this.ctx, key, value).Result()
	if err != nil {
		LOG.Println(err)
		return false, ""
	}
	return true, oldValue
}

// Incr key值每次加一 并返回新值
func (this *RedisClient) Incr(key string) int64 {
	val, err := this.client.Incr(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// IncrBy key值每次加指定数值 并返回新值
func (this *RedisClient) IncrBy(key string, incr int64) int64 {
	val, err := this.client.IncrBy(this.ctx, key, incr).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// IncrByFloat key值每次加指定浮点型数值 并返回新值
func (this *RedisClient) IncrByFloat(key string, incrFloat float64) float64 {
	val, err := this.client.IncrByFloat(this.ctx, key, incrFloat).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// Decr key值每次递减 1 并返回新值
func (this *RedisClient) Decr(key string) int64 {
	val, err := this.client.Decr(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// DecrBy key值每次递减指定数值 并返回新值
func (this *RedisClient) DecrBy(key string, incr int64) int64 {
	val, err := this.client.DecrBy(this.ctx, key, incr).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// Del 删除 key
func (this *RedisClient) Del(key string) bool {
	result, err := this.client.Del(this.ctx, key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// Expire 设置 key的过期时间
func (this *RedisClient) Expire(key string, ex time.Duration) bool {
	result, err := this.client.Expire(this.ctx, key, ex).Result()
	if err != nil {
		return false
	}
	return result
}

/*------------------------------------ list 操作 ------------------------------------*/

// LPush 从列表左边插入数据，并返回列表长度
func (this *RedisClient) LPush(key string, date ...interface{}) int64 {
	result, err := this.client.LPush(this.ctx, key, date).Result()
	if err != nil {
		LOG.Println(err)
	}
	return result
}

// RPush 从列表右边插入数据，并返回列表长度
func (this *RedisClient) RPush(key string, date ...interface{}) int64 {
	result, err := this.client.RPush(this.ctx, key, date).Result()
	if err != nil {
		LOG.Println(err)
	}
	return result
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (this *RedisClient) LPop(key string) (bool, string) {
	val, err := this.client.LPop(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
		return false, ""
	}
	return true, val
}

// RPop 从列表右边删除第一个数据，并返回删除的数据
func (this *RedisClient) RPop(key string) (bool, string) {
	val, err := this.client.RPop(this.ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, val
}

// LIndex 根据索引坐标，查询列表中的数据
func (this *RedisClient) LIndex(key string, index int64) (bool, string) {
	val, err := this.client.LIndex(this.ctx, key, index).Result()
	if err != nil {
		LOG.Println(err)
		return false, ""
	}
	return true, val
}

// LLen 返回列表长度
func (this *RedisClient) LLen(key string) int64 {
	val, err := this.client.LLen(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// LRange 返回列表的一个范围内的数据，也可以返回全部数据
func (this *RedisClient) LRange(key string, start, stop int64) []string {
	vales, err := this.client.LRange(this.ctx, key, start, stop).Result()
	if err != nil {
		LOG.Println(err)
	}
	return vales
}

// LRem 从列表左边开始，删除元素data， 如果出现重复元素，仅删除 count次
func (this *RedisClient) LRem(key string, count int64, data interface{}) bool {
	_, err := this.client.LRem(this.ctx, key, count, data).Result()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

// LInsert 在列表中 pivot 元素的后面插入 data
func (this *RedisClient) LInsert(key string, pivot int64, data interface{}) bool {
	err := this.client.LInsert(this.ctx, key, "after", pivot, data).Err()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return true
}

/*------------------------------------ set 操作 ------------------------------------*/

// SAdd 添加元素到集合中
func (this *RedisClient) SAdd(key string, data ...interface{}) bool {
	err := this.client.SAdd(this.ctx, key, data).Err()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return true
}

// SCard 获取集合元素个数
func (this *RedisClient) SCard(key string) int64 {
	size, err := this.client.SCard(this.ctx, "key").Result()
	if err != nil {
		LOG.Println(err)
	}
	return size
}

// SIsMember 判断元素是否在集合中
func (this *RedisClient) SIsMember(key string, data interface{}) bool {
	ok, err := this.client.SIsMember(this.ctx, key, data).Result()
	if err != nil {
		LOG.Println(err)
	}
	return ok
}

// SMembers 获取集合所有元素
func (this *RedisClient) SMembers(key string) []string {
	es, err := this.client.SMembers(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return es
}

// SRem 删除 key集合中的 data元素
func (this *RedisClient) SRem(key string, data ...interface{}) bool {
	_, err := this.client.SRem(this.ctx, key, data).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return true
}

// SPopN 随机返回集合中的 count个元素，并且删除这些元素
func (this *RedisClient) SPopN(key string, count int64) []string {
	vales, err := this.client.SPopN(this.ctx, key, count).Result()
	if err != nil {
		LOG.Println(err)
	}
	return vales
}

/*------------------------------------ hash 操作 ------------------------------------*/

// HSet 根据 key和 field字段设置，field字段的值
func (this *RedisClient) HSet(key, field, value string) bool {
	err := this.client.HSet(this.ctx, key, field, value).Err()
	if err != nil {
		return false
	}
	return true
}

// HGet 根据 key和 field字段，查询field字段的值
func (this *RedisClient) HGet(key, field string) string {
	val, err := this.client.HGet(this.ctx, key, field).Result()
	if err != nil {
		LOG.Println(err)
	}
	return val
}

// HMGet 根据key和多个字段名，批量查询多个 hash字段值
func (this *RedisClient) HMGet(key string, fields ...string) []interface{} {
	vales, err := this.client.HMGet(this.ctx, key, fields...).Result()
	if err != nil {
		panic(err)
	}
	return vales
}

// HGetAll 根据 key查询所有字段和值
func (this *RedisClient) HGetAll(key string) map[string]string {
	data, err := this.client.HGetAll(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return data
}

// HKeys 根据 key返回所有字段名
func (this *RedisClient) HKeys(key string) []string {
	fields, err := this.client.HKeys(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return fields
}

// HLen 根据 key，查询hash的字段数量
func (this *RedisClient) HLen(key string) int64 {
	size, err := this.client.HLen(this.ctx, key).Result()
	if err != nil {
		LOG.Println(err)
	}
	return size
}

// HMSet 根据 key和多个字段名和字段值，批量设置 hash字段值
func (this *RedisClient) HMSet(key string, data map[string]interface{}) bool {
	result, err := this.client.HMSet(this.ctx, key, data).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return result
}

// HSetNX 如果 field字段不存在，则设置 hash字段值
func (this *RedisClient) HSetNX(key, field string, value interface{}) bool {
	result, err := this.client.HSetNX(this.ctx, key, field, value).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return result
}

// HDel 根据 key和字段名，删除 hash字段，支持批量删除
func (this *RedisClient) HDel(key string, fields ...string) bool {
	_, err := this.client.HDel(this.ctx, key, fields...).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return true
}

// HExists 检测 hash字段名是否存在
func (this *RedisClient) HExists(key, field string) bool {
	result, err := this.client.HExists(this.ctx, key, field).Result()
	if err != nil {
		LOG.Println(err)
		return false
	}
	return result
}
