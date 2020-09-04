package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s1 []int = primes[1:4]
	var s2 []int = primes[1:4]
	s2[0] =1
	fmt.Println(s1,s2)
}