package main

import "fmt"

func main() {
	fmt.Println(f0()) //6
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}


// golang语言中的return语句不是原子操作，
// 分为(1)返回值赋值和(2)RET指令两步。而defer语句执行在赋值之后，RET之前。
// return语句执行步骤
// 1、返回值赋值
// 2、defer语句
// 3、真正RET返回

// defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体。

// 多个defer 后进先出
func f0() (x int) {
	x = 5
	defer func() {
	    fmt.Println("defer_1");
		x++
	}()
	defer func() {
	    fmt.Println("defer_2");
	}()
// 	return x //返回值RET=x, x++, RET=x=6
	return 
}

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x //返回值RET=5, x++, RET=5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值RET=x=5, x++, RET=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值RET=y=x=5, x++, RET=5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //返回值RET=x=5, x`++, RET=5
}


