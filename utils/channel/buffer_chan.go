package main

import "fmt"

func main() {
    // 这次的channel是一个size=1的bufferd channel，将数据发送给channel后，数据被拷贝到channel的缓冲区，goruntine继续往后执行，所以程序可以正常工作
// 	var ch = make(chan string, 1)
// 	ch <- "hello"
// 	fmt.Println(<-ch) //hello
	
// 	当我们调换发送和接收的顺序时，程序又进入了死锁。因为当channel里没有数据时（干涸），执行receive的goroutine也会被block住，最终导致了死锁
	var ch = make(chan string, 1)
	fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan receive]:
	ch <- "hello"
	
// 	对于buffer size=1的channel，第二个数据发送完成之前，之前发送的第一个数据一定被取走了，否则发送也会被block住
// 对于buffer size>1的channel，发送数据时，之前发送的数据不能保证一定被取走了，并且buffer size越大，数据的交付得到的保证越少
// 根据发送数据、接收数据、数据处理的速度合理的设计buffer size，甚至可以在不浪费空间的情况下做到没有任何延迟
// 如果channel buffer已经塞满了数据，继续执行发送会导致当前goruntine被block住（阻塞），直到channel中的数据被取走一部分才可以继续向channel发送数据
}