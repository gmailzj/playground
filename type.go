package main

import "fmt"
import "math"


type person struct {
    age int 
    name string
}

type  myint uint32

// var p person
// func flushICache(begin, end uintptr)

var gvar int  

func main() {
	ages:= map[string]int{"zhangsan": 20}
	fmt.Println(ages)
	modify(ages)
	fmt.Println(ages)
	
// 	jim := person{10,"Jim"}
	// 使用冒号:分开字段名和字段值即可
	jim := person{name:"Jim",age:10}
	fmt.Println(jim)
	modify2(jim)
	fmt.Println(jim)
	
	fmt.Print(math.Floor(5.1))
}

// value
func modify(m map[string]int) {
	m["zhangsan"] = 10
}

// refrence
func modify2(p person) {
	p.age = p.age+10
}
