package main

import "fmt"

// 只在乎行为，不在乎其值
type I interface {
    Get() int
    Set(int)
}

type S struct {
    Age int
}
func(s S) Get() int {
    return s.Age
}
func(s *S) Set(age int) {
    s.Age = age
}

func f(i I){
    i.Set(10)
    fmt.Println(i.Get())
}


// type ByteCounter int
// func (c *ByteCounter) Write(p []byte) (int, error) {
//     *c += ByteCounter(len(p)) // convert int to ByteCounter
//     return len(p), nil
// }
// var c ByteCounter
// c.Write([]byte("hello"))
// fmt.Println(c) // 5  #1

func main() {
	s := S{Age:10} 
    f(&s)  //10
    

    s1 := S{Age:5}
    var i I //声明 i
    i = &s1 //赋值 s 到 i
    fmt.Println(i.Get())
    
    // value, ok := em.(T)：
    // em 是 interface 类型的变量，T代表要断言的类型，
    // value 是 interface 变量存储的值，ok 是 bool 类型表示是否为该断言的类型 T
    if t, ok := i.(*S); ok {
        fmt.Println("s implements I", t)
    }
}