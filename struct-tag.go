package main // abc

import "fmt"
import "reflect"
import "encoding/json"


// struct成员变量标签
type User struct {
    UserId   int    `json:"user_id" bson:"user_id"`
    UserName string `json:"user_name" bson:"user_name"`
}

func main() {
    /*
    var json struct {
		Value string `json:"value" binding:"required"`
	}
	*/
	
	// struct flag
	user := &User{1, "admin"}
    s := reflect.TypeOf(user).Elem() //通过反射获取type定义
    for i := 0; i < s.NumField(); i++ {
        fmt.Println(s.Field(i).Tag) //将tag输出出来
    }
     
    j, _ := json.Marshal(user)
    fmt.Println(string(j))
}