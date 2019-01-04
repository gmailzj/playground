package main

import "fmt"

func main() {
	fmt.Println("BubbleSort");
	arr := []int {4,3,2,1};
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
	BubbleSort(arr);
	fmt.Println(arr)
}

func BubbleSort(values []int) {
    flag := true
    vLen := len(values)
    for i := 0; i < vLen-1; i++ {
        flag = true
        for j := 0; j < vLen-i-1; j++ {
            // 交换  
            if values[j] > values[j+1] {
                values[j], values[j+1] = values[j+1], values[j]
                flag = false
                continue
            }
        }
        if flag {
            break
        }
    }
}
