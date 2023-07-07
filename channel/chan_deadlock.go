package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}
	close(ch) // 这里关闭channel已经”通知“不到range了，会触发死锁。
	// 不管这里是否关闭channel，都会报死锁，close(ch)的位置就不对。
	// 且关闭channel的操作者也错了，只能是发送者关闭channel
}

// 运行程序输出
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// fatal error: all goroutines are asleep - deadlock!
