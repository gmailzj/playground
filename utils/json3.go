package main

import  "fmt"
import "encoding/json"
import "strconv"


type User struct {
    // Golang中，如果指定一个field序列化成JSON的变量名字为-，
    // 则序列化的时候自动忽略这个field
    // Email    string `json:"-"`
    Email    string `json:"email"`
    Password string `json:"password"`
    // many more fields…
}

type User2 struct {
    User
    Password omit `json:"password,omitempty"`
    // Email    string `json:"-"`
    // Token    string `json:"token"`
}

type BlogPost struct {
    URL   string `json:"url"`
    Title string `json:"title"`
}
type Analytics struct {
    Visitors  int `json:"visitors"`
    PageViews int `json:"page_views"`
}

type CacheItem struct {
    Key    string `json:"key"`
    MaxAge int    `json:"cacheAge"`
    Value  string  `json:"cacheValue"`
}
// 因为 bool 默认 为false 
type omit bool

func main() {
    
    
    // 临时忽略struct字段
    user := User{ Email: "gmail@qq.com", Password:"fsafsad"}
    
    user2 := User2{
        User: user,
        // Token: "0xabcdef",
    }
	result, err :=  json.Marshal(user2)
    fmt.Println(err, string(result))
    fmt.Printf("%+v %+v %+v", user, user2, user2.Email)
    
    
    
    // 
    i, err := strconv.ParseInt("123", 10, 64) 
    fmt.Println()
    fmt.Println(i, err)
    
    
    post := &BlogPost{"attila@attilaolah.eu", "Attila's Blog"}
    analytics := &Analytics{Visitors:6, PageViews:4}
    
    r1,_ :=json.Marshal(struct{
        *BlogPost
        *Analytics
    }{post, analytics})
    fmt.Println(string(r1))
    
    
    var post2 BlogPost
    var analytics2 Analytics
    json.Unmarshal([]byte(`{
      "url": "attila@attilaolah.eu",
      "title": "Attila's Blog",
      "visitors": 6,
      "page_views": 14
    }`), &struct {
      *BlogPost
      *Analytics
    }{&post2, &analytics2})
    fmt.Println(post2, analytics2)
    
    // cache item
    item:= &CacheItem{Key:"key1", MaxAge:123}
    
    transfered, _ := json.Marshal(struct{
        *CacheItem
        // Omit bad keys
        OmitMaxAge int  `json:"cacheAge,omitempty"`
        OmitValue  string  `json:"cacheValue,omitempty"`
        // Add nice keys
        MaxAge int    `json:"max_age"`
        Value  string `json:"value"`
    }{
        CacheItem: item,
        // Set the int by value:
        MaxAge: item.MaxAge,
        // Set the nested struct by reference, avoid making a copy:
        Value: item.Value,
    })
    fmt.Println(string(transfered))
}