package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter MyCounter
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 执行10万次累加
			for j := 0; j < 100000; j++ {
				counter.Incr() // 受到锁保护的方法
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count())


}
// 一个线程安全的计数器
type MyCounter struct {
	mu sync.RWMutex
	count uint64
}
// 使用写锁保护
func (c *MyCounter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
// 使用读锁保护
func (c *MyCounter) Count() uint64 {
	c.mu.RLocker()
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
