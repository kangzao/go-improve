package main

import "fmt"

/*
不带缓冲区 channel
*/
func main() {
	ch := make(chan int) // 无缓冲的channel
	go unbufferChan(ch)

	for i := 0; i < 10; i++ {
		fmt.Println("receive ", <-ch) // 读出值
	}
}

func unbufferChan(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ", i)
		ch <- i // 写入值
	}
}
