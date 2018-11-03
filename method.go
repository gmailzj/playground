package main

import (
	"fmt"
	"math"
)

type MyFloat float64


func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 如果不加*就不能set
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
