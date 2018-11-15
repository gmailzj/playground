package main

import "fmt"
import "encoding/json"


var data = []int{1, 2, 3}

type Message struct {
    email string
    phone string
}
// 模拟私有数据
var secrets = map[string] Message{
	"foo": Message{email: "foo@bar.com", phone: "123433", },
}

var list = map[string]Message{
    "foo": Message{"d@qq.com", "abcdef"},
}


type Info struct{
    Name string
    Age int
    Sex string
}

func main() {
    user:="foo"
	fmt.Println("Hello", list["foo"].email)
	if secret, ok := secrets[user]; ok {
	    fmt.Println(user, secret)
	}
	
	//
	var str=[]byte(`[
                {"name":"Sophia","age":23,"sex":"female"},
                {"name":"Benjie","age":32,"sex":"male"}
                    ]`) 
                    
    var info []Info                     
      
    if err:=json.Unmarshal(str,&info);  err!=nil{
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("%v",info) 

}