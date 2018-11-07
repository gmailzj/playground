package main

import "fmt"


func print( s []int32) {
     fmt.Println(s)
}
func main() {
    
    var iarray2  = [5]int32{1, 2, 3, 4, 5}
    iarray3 := [5]int32{1, 2, 3, 4, 5}
    iarray4 := []int32{1, 2, 3, 4, 5}
    fmt.Println(iarray2, iarray3, len(iarray4),cap(iarray4))
    
    iarray5 := iarray3[:]
    print(iarray4);
    fmt.Println(iarray5)
     
    iarray6 := [...]int32{1, 2, 3, 4, 5}
    fmt.Printf("%T,%T,%T, %T", iarray3, iarray4 ,iarray5, iarray6 )

}