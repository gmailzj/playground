package main

import "fmt"
import uuid "github.com/google/uuid"
func main() {
    var a uuid.Time = 64
    uuidMy := uuid.New()
    str := uuidMy.String()
    fmt.Println(str)
	fmt.Println(uuid.New())
	fmt.Printf("%R", a)
	id, err :=uuid.Parse("78fed97d-5505-4cc8-9d98-ee7a45c3627b")
	if(err==nil){
	    fmt.Println(id)
	}
}