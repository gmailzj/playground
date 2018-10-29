package main

import "unsafe"
const (
    a = "abcd"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}