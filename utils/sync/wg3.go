package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	mutexAdd()
}
func mutexAdd() {
	var a int32 =  0
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a += 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use mutex a is %d, spend time: %v\n", a, timeSpends)
}