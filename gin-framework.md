# Gin框架中文文档

### 安装

要安装 Gin 包，你需要已经安装 Go 并且设置好了你的 Go 的工作空间。

 

1. 下载并安装它：

```sh
$ go get -u github.com/gin-gonic/gin
```

2. 在你的代码中导入它：

```go
import "github.com/gin-gonic/gin"
```

3. （可选的） 导入 `net/http` 。 如果你使用常量（如： `http.StatusOK` ） 的时候必须导入。

```go
import "net/http"
```

### 快速开始

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // 在 0.0.0.0:8080 上监听并服务
}
```

### 路由

gin 框架中采用的路由库是 httprouter。

```go
// 创建带有默认中间件的路由:
// 日志与恢复中间件
router := gin.Default()
// 创建不带中间件的路由：
//r := gin.New()
```

支持常见的Http Method

```go
router.GET("/someGet", getting)
router.POST("/somePost", posting)
router.PUT("/somePut", putting)
router.DELETE("/someDelete", deleting)
router.PATCH("/somePatch", patching)
router.HEAD("/someHead", head)
router.OPTIONS("/someOptions", options)
```

##### 获取路由参数

api 参数通过Context的Param方法来获取

```go
router.GET("/user/:name", func(c *gin.Context) {
    // 下面两种方式都可以
    user := c.Params.ByName("name")
    user2 := c.Param("name")
    db["foo"] = "aaa" + user2
    value, ok := db[user]
    if ok {
        c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
    } else {
        c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
    }
})
```

##### 获取URL参数querystring

URL 参数通过 DefaultQuery 或 Query 方法获取

```go
router.POST("post", func(c *gin.Context) {
    // 获取url里面的参数
    id := c.Query("id")
    
    // 如果没有取默认值 取默认值"0"
    page := c.DefaultQuery("page", "0")
    
    message := c.PostForm("message")
    name := c.PostForm("name")
    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
})
```

##### 获取表单参数

表单参数通过 PostForm 方法获取

```go
message := c.PostForm("message")
name := c.PostForm("name")
```

##### Map(URL参数、表单参数)

Map通过QueryMap、PostFormMap 方法获取

```http
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
```

```go
router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
```



##### 路由分组Group

比如 127.0.0.1:8080/v1/user，针对路由已v1/开头的api接口.

```go
v1 := router.Group("/v1")
{
    v1.GET("/user", func(c *gin.Context) {
        c.String(http.StatusOK, "user-v1")
    })
    v1.GET("/list", func(c *gin.Context) {
        c.String(http.StatusOK, "list-v1")
    })

}
```

同时也可以添加中间件，统一控制，比如 Http Basic Authentication

```go
authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":  "bar", // user:foo password:bar
    "manu": "123", // user:manu password:123
}))
```

### 请求和响应

#### 请求

获取请求头

```go
headVersion := context.GetHeader("version")
```

#### 响应

##### 设置响应头

```go
context.Header("lastname", lastname)
```

##### 字符串响应

```go
c.String(http.StatusOK, "some string")
```

##### 格式化json、xml响应

```go
router.GET("/json", func(c *gin.Context) {
		// data := []int{1, 2, 3}
		// c.JSON(http.StatusOK, gin.H{"errCode": 0, "msg": "abc", "data": data})

		var msg struct {
			// 右边的tag用来映射返回结果中的key
			Name    string `json:"user" xml:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
    	// xml
    	// c.XML(http.StatusOK, msg)
	})
```



##### 视图响应

```go
//加载模板
router.LoadHTMLGlob("templates/*")
context.HTML(http.StatusOK, "index.tmpl", gin.H{
    "title": "Main website",
})
```

不同文件夹下模板名字可以相同，此时需要 LoadHTMLGlob() 加载两层模板路径

```go
router.LoadHTMLGlob("templates/**/*")
router.GET("/posts/index", func(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "Posts",
	})
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "Users",
	})
	
}
```

templates/posts/index.tmpl

```go
<!-- 注意开头 define 与结尾 end 不可少 -->
{{ define "posts/index.tmpl" }}
<html><h1>
	{{ .title }}
</h1>
</html>
{{ end }}
```

##### 文件响应

```go
//获取当前文件的相对路径
router.Static("/assets", "./assets")
router.StaticFile("/favicon.ico", "./assets/favicon.ico")
// 没弄明白
router.StaticFS("/static", http.Dir("./assets"))
```

##### 重定向

```go
r.GET("/redirect", func(c *gin.Context) {
	//支持内部和外部的重定向
    c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
})
```



### 中间件

格式：返回一个gin.HandlerFunc的函数

```go
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 在gin上下文中定义变量
		c.Set("example", "12345")

		// 请求前

		c.Next()//处理请求

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
```



#### 分类使用

全局中间件

```go
router.Use(gin.Logger())
router.Use(gin.Recovery())
```

单路由的中间件，可以加任意多个

```
router.GET("/benchmark", MyMiddelware(), benchEndpoint)
```