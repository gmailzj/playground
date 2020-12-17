package main

import "fmt"
import "time"
import "math/rand"
import "strconv"

func main() {
    // 如果channel buffer已经塞满了数据，继续执行发送会导致当前goruntine被block住（阻塞），直到channel中的数据被取走一部分才可以继续向channel发送数据。
	
	const cap = 3
	ch := make(chan string, cap)
 
	for index := 0; index < cap; index++ {
		go func() {
			for p := range ch {
				fmt.Println("Worker received: ", p)
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}()
	}
 
	worknum := 100
	for index := 0; index < worknum; index++ {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		work := "work " + strconv.Itoa(index)
		select {
		case ch <- work:
			fmt.Println("Manager: send a work")
		default:
			fmt.Println("Manager: wait worker")
		}
	}
 
	close(ch)
}