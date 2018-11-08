package main

import "fmt"
import "strings"
import "bytes"

type point struct {
    x, y int
}
func main() {
    var str = "This is a string"
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("%t\n", strings.HasSuffix(str, "ing"))
	
	var hello ="hello"
    var world = "world"
	var buffer bytes.Buffer
        buffer.WriteString(hello)
        buffer.WriteString(",")
        buffer.WriteString(world)
        s := buffer.String()
    fmt.Printf("%v",s)
    
    
    p := point{1, 2}
    fmt.Printf("%v\n", p)

    // If the value is a struct, the `%+v` variant will
    // include the struct's field names.
    fmt.Printf("%+v\n", p)

    // The `%#v` variant prints a Go syntax representation
    // of the value, i.e. the source code snippet that
    // would produce that value.
    fmt.Printf("%#v\n", p)

    // To print the type of a value, use `%T`.
    fmt.Printf("%T\n", p)
    
    // Formatting booleans is straight-forward.
    fmt.Printf("%t\n", true)

    // There are many options for formatting integers.
    // Use `%d` for standard, base-10 formatting.
    fmt.Printf("%d\n", 123)

    // This prints a binary representation.
    fmt.Printf("%b\n", 14)

    // This prints the character corresponding to the
    // given integer.
    fmt.Printf("%c\n", 33)

    // `%x` provides hex encoding.
    fmt.Printf("%x\n", 456)

    // There are also several formatting options for
    // floats. For basic decimal formatting use `%f`.
    fmt.Printf("%f\n", 78.9)
}
