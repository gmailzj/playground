package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice)) //[0 1 2 3 4 5 6 7 8] 9 9
	newSlice := slice[6:8]
	fmt.Println(newSlice, len(newSlice), cap(newSlice)) //[6 7] 2 3
	
	//向newSlice中添加一个元素,注意此时，刚好len(newSlice) == cap(newSlice)
	newSlice = append(newSlice, 111)
	fmt.Println(slice, len(slice), cap(slice))          //[0 1 2 3 4 5 6 7 111] 9 9
	fmt.Println(newSlice, len(newSlice), cap(newSlice)) //[6 7 111] 3 3

	//再向newSlice中添加一个元素,注意此时，
	//刚好向后扩展newSlice时，已经超过了原slice的cap,超过的部分，不会反映到原slice
	newSlice = append(newSlice, 222)
	fmt.Println(slice, len(slice), cap(slice))   //[0 1 2 3 4 5 6 7 111] 9 9
	fmt.Println(newSlice, len(newSlice), cap(newSlice))  //[6 7 111 222] 4 6
}


// package main

// import "fmt"

// func main() {
// 	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println(slice, len(slice), cap(slice)) //[0 1 2 3 4 5 6 7 8] 9 9
// 	newSlice := slice[6:8]
// 	fmt.Println(newSlice, len(newSlice), cap(newSlice)) //[6 7] 2 3

// 	//向newSlice中添加两个元素,注意此时，newSlice的len会变成4，超过了他的cap(3)
// 	//所以会在底层新创建一个数组来保存，所以，对newSlice的操作不会反映到原slice
// 	newSlice = append(newSlice, 111, 222)
// 	fmt.Println(slice, len(slice), cap(slice))          //[0 1 2 3 4 5 6 7 8] 9 9
// 	fmt.Println(newSlice, len(newSlice), cap(newSlice)) //[6 7 111 222] 4 6
// }

// 当我们调用append向一个新的slice添加元素时：

// 　　如果新的slice在append之后，新slice的cap没有改变（即新slice的len仍小于新slice的cap），那么，至少在cap改变之前，append的行为结果都会反映到原slice

// 　　如果新的slice在append之后，新slice的cap改变了，那么，新的slice就拥有了自己的底层数组，所以，append的行为结果不会反映到原slice，但是cap没有改变之前，仍会反映到原slice，只是在cap改变的之后，才不会反映到slice