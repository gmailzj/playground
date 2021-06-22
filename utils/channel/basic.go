package main

import "fmt"
import  "time"
import  "sync"

func main() {
    // eg 1
    // 主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用time.Sleep() 来睡眠一段时间
    // 无法预知for循环内代码运行时间的长短。这时候就不能使用time.Sleep() 来完成等待操作
    for i := 0; i < 100 ; i++{
        go fmt.Println(i)
    }
    time.Sleep(time.Second)
    
    // eg 2
    // 可以考虑使用管道来完成上述操作
    // 但是管道在这里显得有些大材小用，因为它被设计出来不仅仅只是在这里用作简单的同步处理，
    // 在这里使用管道实际上是不合适的。而且假设我们有一万、十万甚至更多的for循环，
    // 也要申请同样数量大小的管道出来，对内存也是不小的开销。
    c := make(chan bool, 100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            c <- true
        }(i)
    }

    for i := 0; i < 100; i++ {
        <-c
    }
    
    // eg 3
    // 使用WaitGroup 将上述代码可以修改为：
    wg := sync.WaitGroup{}
    wg.Add(100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}