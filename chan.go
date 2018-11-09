package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 错误 缓冲区 满了  
    // ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// error  is empty
	fmt.Println(<-ch)

}