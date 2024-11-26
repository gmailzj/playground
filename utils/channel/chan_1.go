package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界1")
	//对于unbufferd channel，不存储任何数据，只负责数据的流通，并且数据的接收一定发生在数据发送完成之前。更详细的解释是，goroutine A在往channel发送数据完成之前，一定有goroutine //B在等着从这个channel接收数据，否则发送就会导致发送的goruntine被block住，所以发送和接收的goruntine是耦合的。

	// 看下面这个例子，往ch发送数据时就使main gouruntine被永久block住，导致程序死锁。
	// 	var ch = make(chan string)
	// 	ch <- "hello" //fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan send]:
	// 	fmt.Println(<-ch)

	// 	有人可能会考虑将接收操作放到前面，不幸的是仍然导致了死锁，
	//  因为channel里没有数据，当前goruntine也会被block住，导致程序死锁。

	//  var ch = make(chan string)
	// 	fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan receive]:
	// 	ch <- "hello"

	// 在另一个goruntine中执行receive，程序就可以正常工作了。
	// 因为在main goruntine发送完数据之前，sub goroutine已经在等待接收数据。

	var ch = make(chan string)
	go func() {
		fmt.Println("Hello, 世界2")
		fmt.Println(<-ch) //out: hello
	}()
	fmt.Println("Hello, 世界3")
	ch <- "hello"
	fmt.Println("Hello, 世界4")
	fmt.Println(time.Millisecond)
}
