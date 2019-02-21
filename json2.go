package main

import "fmt"
import "encoding/json"

type H map[string]interface{}
type I int;
func main() {
	list1 := H{
	  "id":1,
	}
	
	list2 := H{
	  "id":2,
	}
	
	list := []H{
	    list1,
	    list2,
	}
	
	listData := []H{

	}
// 	for _, v := range list {
// 	    v["idstr"] = v["id"]
// 		fmt.Printf("%v\n", v)
// 	}
	for _, v := range list {
	    item := H{}
	    item["id"] = v["id"]
	    item["idstr"] = v["id"]
	    listData = append(listData, item)
// 		fmt.Printf("%v\n", item)
	}
	
	data := H{
	    "id": 1,
	    "list": listData,
	}
	fmt.Println(data)
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))

}