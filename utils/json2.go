package main

// 导入需要的库
import (
    "encoding/json"
    "fmt"
)

// 结构体定义
type Student struct {
    Name    string
    Age     int
    Guake   bool
    Classes []string
    Price   float32
}

// 显示结构体数据，主要是测试和调试的时候需要
func (s *Student) ShowStu() {
    fmt.Println("show Student :")
    fmt.Println("\tName\t:", s.Name)
    fmt.Println("\tAge\t:", s.Age)
    fmt.Println("\tGuake\t:", s.Guake)
    fmt.Println("\tPrice\t:", s.Price)
    fmt.Printf("\tClasses\t: ")
    for _, a := range s.Classes {
        fmt.Printf("%s ", a)
    }
    fmt.Println("")
}

// 主要的函数的调用
func main() {
    st := &Student{
        "Xiao Ming",
        16,
        true,
        []string{"Math", "English", "Chinese"},
        9.99,
    }
    fmt.Println("before JSON encoding :")
    // 打印出结构体的数据的形式
    st.ShowStu()
    // 数据格式打包成josn
    b, err := json.Marshal(st)
    if err != nil {
        fmt.Println("encoding faild")
    } else {
        fmt.Println("encoded data : ")
        fmt.Println(b)
        fmt.Println(string(b))
    }
    // 获取数据的网络格式
    ch := make(chan string, 1)
    go func(c chan string, str string) {
        c <- str
    }(ch, string(b))
    strData := <-ch
    fmt.Println("--------------------------------")

    // 通过josn 转化为结构体格式
    stb := &Student{}
    stb.ShowStu()
    err = json.Unmarshal([]byte(strData), &stb)
    if err != nil {
        fmt.Println("Unmarshal faild")
    } else {
        fmt.Println("Unmarshal success")
        stb.ShowStu()
    }
}