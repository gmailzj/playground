package main

import "fmt"
import "strconv"

func main() {
	
	var i = 65;
	
	// to ascii A  ERROR
	str0 := string(i)
	
    // 通过Itoa方法转换
	str1 := strconv.Itoa(i)
 
	// 通过Sprintf方法转换
	str2 := fmt.Sprintf("%d", i)
    fmt.Println(str0, str1, str2)
    
    // A 65 65
}