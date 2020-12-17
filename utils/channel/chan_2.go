package main

import "fmt"
//import "time"

func main() {
// 	go func() {
//         for i := 0; i < 10; i++ { 
//             fmt.Printf("%d ", i)
//         }
//     }()
//     fmt.Println("Hello, 世界3")
//     fmt.Println(time.Millisecond)
    // 我们期望在sub goruntine中打印10个数，实际上却只有main goruntine打印了hello。
    // 因为在sub goruntine打印之前，main goruntine就已经执行完成并退出了
    // dd := (time.Duration(1) * time.Millisecond)
    // time.Sleep(dd)
    
    
    
    
    
    var ch = make(chan string) 
    go func() {
        for i := 0; i < 10; i++ { 
            fmt.Printf("%d ", i)
        }
        ch <- "exit" 
    }()
    fmt.Println("hello")
    <-ch
}