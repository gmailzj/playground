package main

import "fmt"

//行数
const LINES int = 10

// 杨辉三角
func ShowYangHuiTriangle() {
    nums := []int{}
    for i := 0; i < LINES; i++ {
        //补空白
        for j := 0; j < (LINES - i); j++ {
            fmt.Print(" ")
        }
        for j := 0; j < (i + 1); j++ {
            var length = len(nums)
            var value int
            if j == 0 || j == i {
                value = 1
            } else {
                value = nums[length-i] + nums[length-i-1]
            }
            nums = append(nums, value)
            fmt.Print(value, " ")
        }
        fmt.Println("")
    }
}

func main() {
    ShowYangHuiTriangle()
}