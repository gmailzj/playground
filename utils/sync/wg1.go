package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var n int32
    var wg sync.WaitGroup
    wg.Add(1000)
    for i := 0; i < 1000; i++ {
        // wg.Add(1)
        go func() {
            // error
            n += 1
            //   atomic.AddInt32(&n, 1)
            wg.Done()
        }()
    }
    wg.Wait()

    fmt.Println(atomic.LoadInt32(&n)) // output:1000
}