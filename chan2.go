package main

import "fmt"

func main() {
    c := make(chan int, 1)
    c <- 10
    close(c)
    
    v,ok := <- c // c=10,ok=true，读取出来一个
    v,ok = <- c // c=0,ok=false，实际上没有读出来
    
    fmt.Println(v, ok)

}