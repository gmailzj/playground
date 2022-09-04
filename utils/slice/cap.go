package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
	
	
	s8:= []int{};
	//定义切片s1
    s1 := []int{1, 2, 3}
    //第一种方式：直接声明变量 用=赋值
    //s2切片和s1引用同一个内存地址
    var s2 = s1
    //第二种方式：copy
    var s3 = make([]int, 3)
    copy(s3, s1)            //使用copy函数将 参数2的元素复制到参数1
    fmt.Println(s1, s2, s3, s8) //都是[1 2 3]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
