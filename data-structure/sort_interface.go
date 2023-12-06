package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

// Len()
func (s StuScores) Len() int {
	return len(s)
}

// Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	//	数据集合实现了这三个方法后，即可调用该包的 Sort() 方法进行排序

	// 打印未排序的 stus 数据
	fmt.Println("Default:\n\t", stus)
	// StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)
}
