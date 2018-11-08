package main

import "fmt"
import (
"regexp"
"unicode"
"math"
"strings"
)

func main() {
    var str = "hello"
	fmt.Println(GetStrLength(str))
	
	fmt.Printf("Fields are: %q", strings.Fields("foo bar  baz"))
	
	// 1 分割字符串
	str1 := "1, 2, 3";
	s1 := strings.Split(str1, ",");
	fmt.Printf("\n%v %T", s1, s1);
	// 2  合并为字符串   s1 [][]byte 
	str2 := strings.Join(s1, ",")
	fmt.Printf("\n%v", str2)
	
	// 也就是说，Split 会将 s 中的 sep 去掉，而 SplitAfter 会保留 sep。
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
    fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
    
    // 带 N 的方法可以通过最后一个参数 n 控制返回的结果中的 slice 中的元素个数，
    // 当 n < 0 时，返回所有的子字符串；
    // 当 n == 0 时，返回的结果是 nil；
    // 当 n > 0 时，表示返回的 slice 中最多只有 n 个元素
    // 最后一个元素不会分割
    fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 0))
    
    // s 中是否以 prefix 开始
    // strings.HasPrefix(s, prefix string) bool
    // strings.HasSuffix(s, suffix string) bool
	
}

// GetStrLength 返回输入的字符串的字数，汉字和中文标点算 1 个字数，英文和其他字符 2 个算 1 个字数，不足 1 个算 1个
func GetStrLength(str string) float64 {
	var total float64
 
	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")
 
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 1
		} else {
			total = total + 0.5
		}
	}
 
	return math.Ceil(total)
}