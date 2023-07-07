package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	countNum := 0

	// 确认辅助变量是否都执行完成
	var wg sync.WaitGroup //一个 WaitGroup 对象可以等待一组协程结束

	// wg 添加数目要和 创建的协程数量保持一致
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done() //worker协程执行结束以后，都要调用 wg.Done()；
			for j := 0; j < 1000; j++ {
				mu.Lock()
				countNum++
				mu.Unlock()
			}
		}()
	}
	wg.Wait() //main协程调用 wg.Wait() 且被block，直到所有worker协程全部执行结束后返回。
	fmt.Printf("countNum: %d", countNum)
}
