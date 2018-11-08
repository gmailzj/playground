package main

import "fmt"

func main() {
    var a interface{}
    // a = 12
    a = "abcd"
    
	value, ok := a.(string)
    if !ok {
        fmt.Println("It's not ok for type string")
        return
    }
    
    fmt.Println("The value is ", value)
    
    var t interface{}
    t = 12
    switch t := t.(type) {
        default:
            fmt.Printf("unexpected type %T", t)       // %T prints whatever type t has
        case bool:
            fmt.Printf("boolean %t\n", t)             // t has type bool
        case int:
            fmt.Printf("integer %d\n", t)             // t has type int
        case *bool:
            fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
        case *int:
            fmt.Printf("pointer to integer %d\n", *t) // t has type *int
        }
        
    
}