package main

import "fmt"
import "errors"

import "github.com/davecgh/go-spew/spew"

type input struct {
    id int
    a rune
}
type person struct {
	Name string
}

func main() {
    in := input{
        id:1,
        a:2,
    }
    i:=0
	s:=""
	fmt.Println("Hello, 世界")
	spew.Dump(i,s,in)
	
	m:=map[int]string{1:"1",2:"2"}
	e:=errors.New("嘿嘿，错误")

    p:=person{Name:"张三"}
    spew.Dump(m,e,p)
}