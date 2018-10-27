package main
import "fmt"
func main() {
    var x, y int = 18, 12 
    result := gcd(x,y)  
    fmt.Printf("x, y 的最大公约数是 : %d",result)
}
func gcd(x,y int) int{
     for y != 0  {     
            x, y = y, x%y 
      }  
    return x
}
