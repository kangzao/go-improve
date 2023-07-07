package main

import "fmt"

// 带缓冲区 channel
func main() {
	ch := make(chan string, 3)
	ch <- "tom"
	ch <- "jimmy"
	ch <- "cate"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
