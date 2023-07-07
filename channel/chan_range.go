package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go testCh(ch)
	for val := range ch { //
		fmt.Println("get val: ", val)
	}
}

func testCh(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

// 运行输出：
// get val:  0
// get val:  1
// get val:  2
// get val:  3
// get val:  4
