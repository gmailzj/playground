package main

import "fmt"
import "reflect"

func main() {
    var  s string = "a"
	fmt.Println(reflect.TypeOf(s).String()) // string
	fmt.Println(reflect.ValueOf(s).Kind())  // string
	var  t  = 'a'
	fmt.Println(reflect.TypeOf(t).String()) // int32
	arr := []string{"foo", "bar", "baz"}
	fmt.Println(reflect.TypeOf(arr), reflect.ValueOf(arr).Kind(), reflect.ValueOf(arr).Index(0).Kind())
    
    v := "hello world"
    fmt.Println(typeof(v)) // string
    
    id := 16
    fmt.Println(typeof2(id)) // int
}


// interface{}做为参数
func typeof(v interface{}) string {
    // fmt.Printf(“%T”)里最终调用的还是reflect.TypeOf()
    return fmt.Sprintf("%T", v)
}

func typeof2(v interface{}) string {
    switch t := v.(type) {
    case int:
        return "int"
    case float64:
        return "float64"
    //... etc
    default:
        _ = t
        return "unknown"
    }
}

