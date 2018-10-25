package main

import (
	"fmt"
	"math/rand"
)
// 1.包名过于复杂或者意思不明确。不确定此 db 是哪种类型，故可以使用别名来明确含义：
// 2.包名和其他包冲突。现在我们有库 db ,但不久 出现了另一种DB，叫云DB。但包名是一样的，分别用别名区分
// import mongo "mywebapp/libs/mongodb/db"
// import ydbgo "mywebapp/libs/yundb/db"

// 这里的点.符号表示，对包 lib 的调用直接省略包名
// import . "ysqi/lib"
// 这里的下划线_表示，不准备现在使用你们家的东西，但得提前告诉一声。你先做好准备
// import _ "ysqi/lib"


func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
