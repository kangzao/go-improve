package main

import (
	"fmt"
)

// chan_for.go  循环channel
func main() {
	ch := make(chan int)
	go testLoop(ch)

	for {
		val, ok := <-ch
		if ok == false { // ok 为 false，没有数据可读
			break // 跳出循环
		}
		fmt.Println("get val: ", val)
	}
}

func testLoop(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
