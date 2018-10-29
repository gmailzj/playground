package main

import "fmt"


//一般函数
func func_name1(a int) {
	println(a)
}

//多参数，无返回值
func func_name2(a, b int, c string) {
	println(a, b, c)
}

//单个返回值
func func_name3(a, b int) int { //同类型，可以省略  a, b int
	return a + b
}

//多个返回值
func func_name4(a, b int) (c int, err error) {  //返回值还可以是   (int, error)
	return a+b, nil
}

func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main() {
	fmt.Println("Hello, 世界")
	test1("a", "b")
	var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
    
    // for循环
    for _, v := range strss {
		fmt.Println(v)
	}
}