package test

import (
	"code-go/middleware"
	"fmt"
	"testing"
	"time"
)

// 在上面的代码中，我们定义了一个 TokenBucket 结构体，表示令牌桶。
// 它包含桶的容量 capacity、当前令牌数量 tokens、每秒生成的令牌数 rate 和上次生成令牌的时间 lastToken。
// 通过 NewTokenBucket 函数创建一个新的令牌桶实例。
//
// 在 Allow 方法中，我们首先计算上次生成令牌到现在的时间间隔内，应该生成的令牌数量。
// 然后，根据生成的令牌数量，更新令牌桶的状态。最后，判断当前令牌桶中是否有足够的令牌可用，如果有则返回 true，否则返回 false。
//
// 在 main 函数中，我们创建了一个容量为 10，每秒生成 5 个令牌的令牌桶。
// 然后，模拟请求的处理，通过调用 Allow 方法判断是否允许处理请求。根据返回结果，可以进行相应的处理。
func TestTokenBucket(t *testing.T) {
	// 创建一个容量为 10，每秒生成 5 个令牌的令牌桶
	tb := middleware.NewTokenBucket(10, 5)

	go func(tb *middleware.TokenBucket) {
		// 模拟请求的处理
		for i := 1; i <= 2000; i++ {
			if tb.Allow() {
				fmt.Println("Request", i, "is allowed")
				// 处理请求
			} else {
				fmt.Println("Request", i, "is denied")
				// 拒绝请求或进行其他处理
			}
			time.Sleep(time.Millisecond * 200)
		}
	}(tb)

	// 模拟请求的处理
	for i := 1; i <= 2000; i++ {
		if tb.Allow() {
			fmt.Println("Request", i, "is allowed")
			// 处理请求
		} else {
			fmt.Println("Request", i, "is denied")
			// 拒绝请求或进行其他处理
		}
		time.Sleep(time.Millisecond * 200)
	}

}
