# 文档

## godoc

golang 官方有有文档自动生成网站，地址是 godoc.org，比如：logger 的[文档](https://godoc.org/github.com/mnhkahn/gogogo/logger)

`godoc`也可以在本地启动：

```bash
godoc -http=:6060
# -play可以使用playground运行Example代码
godoc -http=:6060 -play 
```

## 根据代码注释查看 go doc

```
go doc 
```

*go doc命令的标记说明*

| 标记名称 | 标记描述                                                     |
| :------- | :----------------------------------------------------------- |
| -c       | 加入此标记后会使`go doc`命令区分参数中字母的大小写。默认情况下，命令是大小写不敏感的。 |
| -cmd     | 加入此标记后会使`go doc`命令同时打印出`main`包中的可导出的程序实体（其名称的首字母大写）的文档。默认情况下，这部分文档是不会被打印出来的。 |
| -u       | 加入此标记后会使`go doc`命令同时打印出不可导出的程序实体（其名称的首字母小写）的文档。默认情况下，这部分文档是不会被打印出来的。 |

`go doc`命令可以后跟一个或两个参数。当然，我们也可以不附加任务参数。如果不附加参数，那么`go doc`命令会试图打印出当前目录所代表的代码包的文档及其中的包级程序实体的列表。

```
$ go doc service
package service // import "goods-manage/service"

Package service implements methods for services.

func NewGoodsManageService(ctx context.Context) *goodsManageService
```



## 什么是分布式锁？

分布式锁是控制分布式系统或不同系统之间共同访问共享资源的一种锁实现，如果不同的系统或同一个系统的不同主机之间共享了某个资源时，往往需要互斥来防止彼此干扰来保证一致性。

## 分布式锁需要具备哪些条件

1. 互斥性：在任意一个时刻，只有一个客户端持有锁。
2. 无死锁：即便持有锁的客户端崩溃或者其他意外事件，锁仍然可以被获取。
3. 容错：只要大部分Redis节点都活着，客户端就可以获取和释放锁

## go语言实现

https://github.com/go-redsync/redsync

[http://redis.io/topics/distlock](http://redis.io/topics/distlock)



# Defer

错误处理

```go
defer func() {
		if err := recover(); err != nil {
			logger.Info(logId, "GoodsManageImp:QueryOne:错误", err)
		}
	}()
```

Go 报错 invalid memory address or nil pointer dereference

仔细排查后发现问题出在这个块代码

```go
rows, err := db.Query(query, args...)
defer rows.Close()
if err != nil {
    return nil, err
}
```

经过 Google 后发现，`defer` 的位置不正确

> The defer only defers the function call. The field and method are accessed immediately.

修改后即可

```go
rows, err := db.Query(query, args...)
if err != nil {
    return nil, err
}
defer rows.Close()
```

