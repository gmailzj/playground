package main

import "fmt"
import (
    "strconv"
    "strings"
    "unicode/utf8"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (p IPAddr) String() string {
	var ret = [] string{} 
	for _, value := range p {
		ret = append(ret, fmt.Sprintf("%d", value))
	}
	return fmt.Sprintf("%s", strings.Join(ret, ","))
}

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
    
    hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	
	str := "hello中"
	// RuneCountInString more efficient
	fmt.Println(len(str), len([]rune(str)), utf8.RuneCountInString(str))
}