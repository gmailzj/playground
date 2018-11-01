package main

import "fmt"
import "math/rand"
import "time"
import "math/big"

func init(){
    //以时间作为初始化种子
    rand.Seed(time.Now().UnixNano())
}

func main() {
    freq := rand.Float64() * 5.0
	fmt.Println(freq)
	freq = rand.Float64() * 5.0
	fmt.Println(freq)
	max := big.NewInt(100)
	fmt.Println(max)
	
	
    a := rand.Int()
    fmt.Println(a)

    b := rand.Intn(100)
    fmt.Println(b)

    c := rand.Float32()
    fmt.Println(c)
    
}