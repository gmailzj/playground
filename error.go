package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


func Sqrt(x float64) (float64, error) {
	if(x<0){
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}



func main() {
	v, err := Sqrt((12));
	if  err != nil{
		fmt.Printf("couldn't convert number: %v\n", err)
		return 
	}
	fmt.Println(v)
}
