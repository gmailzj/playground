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

