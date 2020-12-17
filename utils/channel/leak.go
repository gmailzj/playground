package main

import "fmt"
import "time"
import "math/rand"

func main() {
	ch := make(chan string)
	count := 3
	for index := 0; index < count; index++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "hello"
		}()
	}
	
	// 泄露
// 	fmt.Println(<-ch)
	// 改进 
	for index := 0; index < count; index++ {
		fmt.Println(<-ch)
	}
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
}