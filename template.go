package main

import (
    "html/template"
    "os"
)

type Person struct {
    UserName string
    Password    string  //未导出的字段，首字母是小写的
}

func main() {
    t := template.New("tmp")
    
    // 使用Parse代替ParseFiles，因为Parse可以直接测试一个字符串，而不需要额外的文件
    t, _ = t.Parse("<div>hello {{.UserName}} {{.Password}}!</div>")
    p := Person{UserName: "world"}
    // 使用os.Stdout代替http.ResponseWriter，因为os.Stdout实现了io.Writer接口
    t.Execute(os.Stdout, p)
}