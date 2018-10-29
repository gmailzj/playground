package main

import "fmt"

type T struct {
}
func (t *T) String() string {
    return ""
}
func main() {
	var t T
    fmt.Println(t.String())
    
}