package main

import "fmt"
import "strings"

func main() {
    var str = "This is a string"
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

}
