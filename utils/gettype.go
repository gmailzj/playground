package main
import (
    "reflect"
    "fmt"
)
func main() {
    v := "hello world"
    fmt.Println(typeof(v))
    fmt.Println(typeof2(v))
    fmt.Println(typeof3(v))
}

func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}
func typeof2(v interface{}) string {
    return reflect.TypeOf(v).String()
}
func typeof3(v interface{}) string {
    switch t := v.(type) {
    case int:
        return "int"
    case string:
        return "string"
    case float64:
        return "float64"
    //... etc
    default:
        _ = t
        return "unknown"
    }
}