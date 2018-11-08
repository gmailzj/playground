package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	
	s = append(s, 12, 11, 10, 9)
	printSlice(s)
	
	slice := append([]int{1,2,3},4,5,6)
	printSlice(slice)
	s1 := []int{1,2,3}
	s2 := []int {4,5,6}
	
// 	append函数返回值必须有变量接收，不然编译器会报错
	s3 := append(s1, s2...)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
