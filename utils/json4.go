package main

import "fmt"
import "encoding/json"
type A struct {
    name *string
}


type User struct {

    Email    string `json:"email"`
    Password string `json:"password"`

}

type User2 struct {
    *User
    Password string `json:"password,omitempty"`
    Token string `json:"token,omitempty"`
}

func main() {
   user := User2{ User: &User{Email:"s",Password:"true"}, Token:"s"}
   userString,_ := json.Marshal(user);
   fmt.Println(user.Password, string(userString))
}