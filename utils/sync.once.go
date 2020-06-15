package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	onceBody2 := func() {
		fmt.Println("Only once2")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			once.Do(onceBody2)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
// Only once