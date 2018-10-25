package main

import "fmt"

import "reflect"

// 函数可以返回任意数量的返回值。
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b, typeof(a))
    fmt.Println(a, b, typeof2(a))
    fmt.Println(a, b, typeof3(a))
}



func typeof2(v interface{}) string {
    return reflect.TypeOf(v).String()
}

func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func typeof3(v interface{}) string {
    switch t := v.(type) {
    case int:
        return "int"
    case float64:
        return "float64"
    case string:
        return "string"
    //... etc
    default:
        _ = t
        return "unknown"
    }
}
