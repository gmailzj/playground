package main

import "fmt"
import "encoding/json"

// 因为 bool 默认 为false 
type omit bool
type User struct {

    Email    string `json:"email"`
    Password string `json:"password"`

}

type User2 struct {
    *User
    Password bool `json:"password,omitempty"`
    Token string `json:"token,omitempty"`
}

func main() {
   // 临时忽略struct字段
   user := User2{ User: &User{Email:"s",Password:"true"}, Token:"s"}
   userString,_ := json.Marshal(user);
   fmt.Println(user.Password, string(userString))
}