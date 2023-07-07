package main

import "fmt"

func main() {
	//可读可写的 channel
	ch := make(chan int)
	go testData(ch)
	fmt.Println(<-ch)
}

func testData(ch chan<- int) {
	ch <- 10
}

// 运行输出：
// 10

// 没有报错，可以正常输出结果
