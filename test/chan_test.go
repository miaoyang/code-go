package test

import (
	"fmt"
	"testing"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		// 模拟处理任务的耗时操作
		for i := 0; i < 100000000; i++ {
		}
		results <- j * 2
		fmt.Println("worker", id, "finished job", j)
	}
}

func TestChannl(t *testing.T) {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动多个协程处理任务
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 发送任务给协程
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 获取处理结果
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Println("result:", result)
	}
	close(results)
}
