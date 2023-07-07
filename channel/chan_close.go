package main

import (
	"fmt"
)

// 关闭channel
func main() {
	ch := make(chan int)
	go test(ch)

	for {
		if v, ok := <-ch; ok {
			fmt.Println("get val: ", v, ok)
		} else {
			break
		}

	}
}

func test(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
