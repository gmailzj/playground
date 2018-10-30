## 前言

Go 是一门简单有趣的编程语言，与其他语言一样，在使用时不免会遇到很多坑，不过它们大多不是 Go 本身的设计缺陷。如果你刚从其他语言转到 Go，那这篇文章里的坑多半会踩到。

如果花时间学习官方 doc、wiki、[讨论邮件列表](https://groups.google.com/forum/#!forum/golang-nuts)、 [Rob Pike](https://github.com/robpike) 的大量文章以及 Go 的源码，会发现这篇文章中的坑是很常见的，新手跳过这些坑，能减少大量调试代码的时间。

## 初级篇：1-34

### 1. 左大括号 `{` 不能单独放一行

在其他大多数语言中，`{` 的位置你自行决定。Go 比较特别，遵守分号注入规则（automatic semicolon injection）：编译器会在每行代码尾部特定分隔符后加 `;` 来分隔多条语句，比如会在 `)` 后加分号：

```go
// 错误示例
func main()                    
{
    println("hello world")
}
 
// 等效于
func main();    // 无函数体                    
{
    println("hello world")
}
```

```bash
./main.go: missing function body
./main.go: syntax error: unexpected semicolon or newline before {
```

```go
// 正确示例
func main() {
    println("hello world")
}     
```

### 2. 未使用的变量

如果在函数体代码中有未使用的变量，则无法通过编译，不过全局变量声明但不使用是可以的。

即使变量声明后为变量赋值，依旧无法通过编译，需在某处使用它：

```go
// 错误示例
var gvar int     // 全局变量，声明不使用也可以
 
func main() {
    var one int     // error: one declared and not used
    two := 2    // error: two declared and not used
    var three int    // error: three declared and not used
    three = 3        
}
 
 
// 正确示例
// 可以直接注释或移除未使用的变量
func main() {
    var one int
    _ = one
    
    two := 2
    println(two)
    
    var three int
    one = three
 
    var four int
    four = four
}
```

### 3. 未使用的 import

如果你 import 一个包，但包中的变量、函数、接口和结构体一个都没有用到的话，将编译失败。

可以使用 `_` 下划线符号作为别名来忽略导入的包，从而避免编译错误，这只会执行 package 的 `init()`

```go
// 错误示例
import (
    "fmt"    // imported and not used: "fmt"
    "log"    // imported and not used: "log"
    "time"    // imported and not used: "time"
)
 
func main() {
}
 
 
// 正确示例
// 可以使用 goimports 工具来注释或移除未使用到的包
import (
    _ "fmt"
    "log"
    "time"
)
 
func main() {
    _ = log.Println
    _ = time.Now
}
```

### 4. 简短声明的变量只能在函数内部使用

```go
// 错误示例
myvar := 1    // syntax error: non-declaration statement outside function body
func main() {
}
 
 
// 正确示例
var  myvar = 1
func main() {
}
```

### 5. 使用简短声明来重复声明变量

不能用简短声明方式来单独为一个变量重复声明， `:=` 左侧至少有一个新变量，才允许多变量的重复声明：

```go
// 错误示例
func main() {  
    one := 0
    one := 1 // error: no new variables on left side of :=
}
 
 
// 正确示例
func main() {
    one := 0
    one, two := 1, 2    // two 是新变量，允许 one 的重复声明。比如 error 处理经常用同名变量 err
    one, two = two, one    // 交换两个变量值的简写
}
```

### 6. 不能使用简短声明来设置字段的值

struct 的变量字段不能使用 `:=` 来赋值以使用预定义的变量来避免解决：

```go
// 错误示例
type info struct {
    result int
}

func work() (int, error) {
    return 3, nil
}

func main() {
    var data info
    data.result, err := work()    // error: non-name data.result on left side of :=
    fmt.Printf("info: %+v\n", data)
}
```

```go
// 正确示例
func main() {
    var data info
    var err error    // err 需要预声明
    
    data.result, err = work()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("info: %+v\n", data)
}
```

### 7. 不小心覆盖了变量

对从动态语言转过来的开发者来说，简短声明很好用，这可能会让人误会 `:=` 是一个赋值操作符。

如果你在新的代码块中像下边这样误用了 `:=`，编译不会报错，但是变量不会按你的预期工作：

```go
func main() {
    x := 1
    println(x)        // 1
    {
        println(x)    // 1
        x := 2
        println(x)    // 2    // 新的 x 变量的作用域只在代码块内部
    }
    println(x)        // 1
}
```

可使用 [vet](http://godoc.org/golang.org/x/tools/cmd/vet) 工具来诊断这种变量覆盖，Go 默认不做覆盖检查，添加 `-shadow` 选项来启用：

> go tool vet -shadow main.go
> main.go:9: declaration of "x" shadows declaration at main.go:5

注意 vet 不会报告全部被覆盖的变量，可以使用 [go-nyet](https://github.com/barakmich/go-nyet) 来做进一步的检测：

> $GOPATH/bin/go-nyet main.go
> main.go:10:3:Shadowing variable `x`

### 8. 显式类型的变量无法使用 nil 来初始化

`nil` 是 interface、function、pointer、map、slice 和 channel 类型变量的默认初始值。但声明时不指定类型，编译器也无法推断出变量的具体类型。

```go
// 错误示例
func main() {
    var x = nil    // error: use of untyped nil
    _ = x
}

// 正确示例
func main() {
    var x interface{} = nil
    _ = x
}    
```


