# 日志处理

beego 的日志处理是基于 logs 模块搭建的，内置了一个变量 `BeeLogger`，默认已经是 `logs.BeeLogger` 类型，初始化了 console，也就是默认输出到 `console`。

## 使用入门

一般在程序中我们使用如下的方式进行输出：

```go
beego.Emergency("this is emergency")
beego.Alert("this is alert")
beego.Critical("this is critical")
beego.Error("this is error")
beego.Warning("this is warning")
beego.Notice("this is notice")
beego.Informational("this is informational")
beego.Debug("this is debug")
```

## 设置输出

我们的程序往往期望把信息输出到 log 中，现在设置输出到文件很方便，如下所示：

```
beego.SetLogger("file", `{"filename":"logs/test.log"}`)
```

更多详细的日志配置请查看 [日志配置](https://beego.me/docs/module/logs.md)

这个默认情况就会同时输出到两个地方，一个 console，一个 file，如果只想输出到文件，就需要调用删除操作：

```
beego.BeeLogger.DelLogger("console")
```

## 设置级别

日志的级别如上所示的代码这样分为八个级别：

```
LevelEmergency
LevelAlert
LevelCritical
LevelError
LevelWarning
LevelNotice
LevelInformational
LevelDebug
```

级别依次降低，默认全部打印，但是一般我们在部署环境，可以通过设置级别设置日志级别：

```
beego.SetLevel(beego.LevelInformational)
```

## 输出文件名和行号

日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置

```
beego.SetLogFuncCall(true)
```

开启传入参数 true, 关闭传入参数 false, 默认是关闭的.