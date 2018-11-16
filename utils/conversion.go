package main

import (
    "log"
    "strconv"
    "reflect"
)

func main() {
    // 进制conversion 
    num := int64(100)
    for i := 2; i <= 36; i++ {
		str := strconv.FormatInt(num, i)
		log.Println("str = ", str)
	}
	num2 := uint64(9)
	// return []byte
	numbin := strconv.AppendUint(nil, num2, 2)
	log.Println("numbin = ", numbin)
	log.Println("num2 = ", num2, " binary = ", string(numbin))
	log.Println(reflect.ValueOf(numbin).Index(0).Kind()) //uint8
}