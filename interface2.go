package main

import "fmt"

type Reader interface {
    Read(p []byte)(n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}

// 组合interface 
// 通过这种形式可以便捷的实现一个新的 interface 而不必写出 interface 包含的所有 method

type ReadWriter interface {
    Reader
    Writer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// []interface{} 其长度是固定的N*2，但是 []T 的长度是N*sizeof(T)
func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}
func main() {
	fmt.Println("Hello, 世界")
	names := []string{"stanley", "david", "oscar"}
// 	cannot use names (type []string) as type []interface {} in argument to printAll
// 	printAll(names)
}