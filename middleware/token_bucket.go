package middleware

import "time"

type TokenBucket struct {
	capacity  int       // 桶的容量
	tokens    int       // 当前令牌数量
	rate      int       // 每秒生成的令牌数
	lastToken time.Time // 上次生成令牌的时间
}

func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		tokens:    capacity,
		rate:      rate,
		lastToken: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	now := time.Now()
	elapsed := now.Sub(tb.lastToken)
	// 计算上次生成令牌到现在的时间间隔内，应该生成的令牌数量
	generatedTokens := int(elapsed.Seconds()) * tb.rate
	if generatedTokens > 0 {
		// 根据生成的令牌数量，更新令牌桶状态
		tb.tokens = tb.tokens + generatedTokens
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastToken = now
	}
	// 判断当前令牌桶中是否有足够的令牌可用
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
