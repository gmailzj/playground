

Go 的并发属于 CSP 并发模型的一种实现，CSP 并发模型的核心概念是：“不要通过共享内存来通信，而应该通 
过通信来共享内存”。这在 Go 语言中的实现就是 Goroutine 和 Channel。

## 在线书籍教程

[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md)

[Go语言教程](http://www.runoob.com/go/go-data-types.html)

https://gobyexample.com/regular-expressions

**go中文文档**

https://go-zh.org/pkg/sync/

## Go 命令基础

### go env

| 名称          | 说明                   |
| ----------- | -------------------- |
| CGO_ENABLED | 指明cgo工具是否可用的标识。      |
| GOARCH      | 程序构建环境的目标计算架构。       |
| GOBIN       | 存放可执行文件的目录的绝对路径。     |
| GOCHAR      | 程序构建环境的目标计算架构的单字符标识。 |
| GOEXE       | 可执行文件的后缀。            |
| GOHOSTARCH  | 程序运行环境的目标计算架构。       |
| GOOS        | 程序构建环境的目标操作系统。       |
| GOHOSTOS    | 程序运行环境的目标操作系统。       |
| GOPATH      | 工作区目录的绝对路径。          |
| GORACE      | 用于数据竞争检测的相关选项。       |
| GOROOT      | Go语言的安装目录的绝对路径。      |
| GOTOOLDIR   | Go工具目录的绝对路径。         |

### go build

go build 命令用于编译我们指定的源码文件或代码包以及它们的依赖包。

Go语言的源码文件有三大类，即：命令源码文件、库源码文件和测试源码文件。他们的功用各不相同，而写法也各有各的特点。命令源码文件总是作为可执行的程序的入口。库源码文件一般用于集中放置各种待被使用的程序实体（全局常量、全局变量、接口、结构体、函数等等）。而测试源码文件主要用于对前两种源码文件中的程序实体的功能和性能进行测试。另外，后者也可以用于展现前两者中程序的使用方法

让Go语言编译器自行决定可执行文件的名称，我们还可以自定义它 -o 

go build -o fn -v main.go 

在执行命令时加入一个标记`-v`。这个标记的意义在于可以使命令把执行过程中构建的包名打印出来

| 标记名称  | 标记描述                                     |
| ----- | ---------------------------------------- |
| -a    | 强行对所有涉及到的代码包（包含标准库中的代码包）进行重新构建，即使它们已经是最新的了。 |
| -n    | 打印编译期间所用到的其它命令，但是并不真正执行它们。               |
| -p n  | 指定编译过程中执行各任务的并行数量（确切地说应该是并发数量）。在默认情况下，该数量等于CPU的逻辑核数。但是在`darwin/arm`平台（即iPhone和iPad所用的平台）下，该数量默认是`1`。 |
| -race | 开启竞态条件的检测。不过此标记目前仅在`linux/amd64`、`freebsd/amd64`、`darwin/amd64`和`windows/amd64`平台下受到支持。 |
| -v    | 打印出那些被编译的代码包的名字。                         |
| -work | 打印出编译时生成的临时工作目录的路径，并在编译结束时保留它。在默认情况下，编译结束时会删除该目录。 |
| -x    | 打印编译期间所用到的其它命令。注意它与`-n`标记的区别。            |

### go install

命令`go install`用于编译并安装指定的代码包及它们的依赖包

`go install`命令只比`go build`命令多做了一件事，即：安装编译后的结果文件到指定目录。

#### go mod

git config --global url."git@gitlab.xxx.com:".insteadOf "http://gitlab.xxx.com/" 

export http_proxy=http://127.0.0.1:1087;   export https_proxy=http://127.0.0.1:1087;

export GOPROXY="https://athens.azurefd.net"
export GO111MODULE=on

## Go 语言最主要的特性：

- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

## 语言的描述——BNF范式

乔姆斯基(Chomsky)于1956年建立形式语言的描述以来，形式语言的理论发展很快。他将[文法分成四种类型](http://baike.baidu.com/link?url=azwDk7HZGexx5dH2iagrxMyA3pXIyqyc5nY50DvVlze13rzGAQ8hoXnwE8tmUbwmwy-CJPkyDv94uTKyRB0uj_)，即0型、1型、2型和3型，0型即自然语言文法，1型称为上下文相关文法，2型称为上下文无关文法，3型称为正则文法。

BNF范式是一种用于表示[上下文无关文法](https://zh.wikipedia.org/wiki/%E4%B8%8A%E4%B8%8B%E6%96%87%E6%97%A0%E5%85%B3%E6%96%87%E6%B3%95)的语言

**巴克斯范式的内容：**

  尖括号( < > )内包含的为必选项。

  : : = 是 “被定义为”的意思。

  方括号( [ ] )内包含的为可选项。

  大括号( { } )内包含的为可重复0至无数次的项。

 竖线( | )表示在其左右两边任选一项，相当于"OR"的意思。

省略号(...)表示该元素可以重复任意多次。如果省略号后面出现分组元素，重复括号里面指定的分组元素。如果省略后出现一个单元素的元素，只是重复单元素。



有了BNF范式定义我们的文法规则，那么我们在识别一个文法的时候，就能对其构建一棵**抽象语法树（AST）**，这个语法树就是帮助我们进行代码编译工作的核心数据结构。

Go语言规范 [The Go Programming Language Specification](https://golang.org/ref/spec#defer_statements)

The syntax is specified using Extended Backus-Naur Form (EBNF):

```
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lower-case production names are used to identify lexical tokens. Non-terminals are in CamelCase. Lexical tokens are enclosed in double quotes `""` or back quotes ````.

The form `a … b` represents the set of characters from `a` through `b` as alternatives. The horizontal ellipsis `…` is also used elsewhere in the spec to informally denote various enumerations or code snippets that are not further specified. The character `…` (as opposed to the three characters `...`) is not a token of the Go language.

## GO 语法基础

### 文件名

Go 的源文件以 `.go` 为后缀名存储在计算机中，这些文件名均由小写字母组成，如 `scanner.go` 。如果文件名由多个部分组成，则使用下划线 `_` 对它们进行分隔，如 `scanner_test.go` 。文件名不包含空格或其他特殊字符。

Go 语言的基础组成有以下几个部分：

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

程序中可能会使用到这些标点符号：.   ,   ;  : 和 …。

### 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

以下是有效的标识符：

```go
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal
```

以下是无效的标识符：

- 1ab（以数字开头）
- case（Go 语言的关键字）
- a+b（运算符是不允许的）

`_` 本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个这个标识符作为变量对其它变量的进行赋值或运算。

在编码过程中，你可能会遇到没有名称的变量、类型或方法。虽然这不是必须的，但有时候这样做可以极大地增强代码的灵活性，这些变量被统称为匿名变量。

### 关键字

下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |



## Go语言数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

Go 语言按类别有以下几种数据类型：

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **布尔型** 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
| 2    | **数字类型** 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型:** 字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本。 |
| 4    | **派生类型:** 包括：(a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型 |

###  默认值

 整形如int8、byte、int16、uint、uintprt等，默认值为0。
 浮点类型如float32、float64，默认值为0。
 布尔类型bool的默认值为false。
 复数类型如complex64、complex128，默认值为0+0i。
 字符串string的默认值为”“。
 错误类型error的默认值为nil。
 对于一些复合类型，如指针、切片、字典、通道、接口，默认值为nil。而数组的默认值要根据其数据类型来确定。例如：var a [4]int，其默认值为[0 0 0 0]

### 数字类型

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

| 序号   | 类型和描述                                    |
| ---- | ---------------------------------------- |
| 1    | **uint8**无符号 8 位整型 (0 到 255)             |
| 2    | **uint16**无符号 16 位整型 (0 到 65535)         |
| 3    | **uint32**无符号 32 位整型 (0 到 4294967295)    |
| 4    | **uint64**无符号 64 位整型 (0 到 18446744073709551615) |
| 5    | **int8**有符号 8 位整型 (-128 到 127)           |
| 6    | **int16**有符号 16 位整型 (-32768 到 32767)     |
| 7    | **int32**有符号 32 位整型 (-2147483648 到 2147483647) |
| 8    | **int64**有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

| 序号   | 类型和描述                       |
| ---- | --------------------------- |
| 1    | **float32**IEEE-754 32位浮点型数 |
| 2    | **float64**IEEE-754 64位浮点型数 |
| 3    | **complex64**32 位实数和虚数      |
| 4    | **complex128**64 位实数和虚数     |

| 序号   | 类型和描述                                    |
| ---- | ---------------------------------------- |
| 1    | **byte**类似 uint8                         |
| 2    | **rune**类似 int32  *Go*语言将*rune*定义为int32类型的别名 |
| 3    | **uint**32 或 64 位                        |
| 4    | **int**与 uint 一样大小                       |
| 5    | **uintptr**无符号整型，用于存放一个指针                |

### 格式化约定

fmt.Printf("%格式")

| 占位符 | 说明                               |
| ------ | ---------------------------------- |
| %v     | 值的默认格式表示                   |
| %+v    | 类似%v，但输出结构体时会添加字段名 |
| %#v    | 值的golang语法表示                 |
| %T     | 打印值的类型                       |
| %%     | 百分号                             |
| %d    | 10进制数,如果d前面有数字，表示控制输出宽度，默认空白填充，%05d 会在不满5位时填充0 |
| %b    | 表示2进制数            |
| %o | 打印整型数，八进制输出，如果o前面带有#表示带有零的八进制 |
| %x | 打印整型数，十六进制输出，如果x前面带有#表示带有0x的十六进制 |
| %f    | 浮点数，有小数         |
| %9.2f | 宽度9，精度2,如果f前面有x.y 表示控制宽度和输出小数点位数，要对齐，再加-，例如 %-10.5f |
| %e    | 科学计数法             |
| %s    | 打印字符串,直接输出字符串或[]byte |
| %q   | 打印字符串,该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %t   |打印布尔值|
| %p   | 指针，表示未16进制，并加上前缀0x                             |

```go
%v 打印结构体
%+v 打印带有字段的结构体
%T 打印对象类型
%t 打印布尔值
%d 打印整型数，十进制输出，
如果d前面有数字，表示控制输出宽度，默认使用空白填充，%05d 会在不满5位时填充0
%b 打印整型数，二进制输出
%c 打印整型数，字符输出（如果有）
%o 打印整型数，八进制输出，如果o前面带有#表示带有零的八进制
%x 打印整型数，十六进制输出，如果x前面带有#表示带有0x的十六进制
%f 打印浮点数，正常输出，
如果f前面有x.y 表示控制宽度和输出小数点位数，要对其，再加-，例如 %-10.5f
%e,%E 打印浮点数，科学记数法输出
%s 打印字符串
%q 打印字符串，带有引号输出
%x 打印字符串，使用base-16输出其编码
%p 打印指针
%U 打印Unicode字符
%#U 打印带字符的Unicode
```

General:

```go
%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value; 相应值的 Go 语法表示
%T	a Go-syntax representation of the type of the value; 相应值的类型的 Go 语法表示
```

The default format for %v is:

```go
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

String and slice of bytes (treated equivalently with these verbs):

```go
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte
```

Boolean:

```go
%t	the word true or false
```

Integer:

```go
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
```

Floating-point and complex constituents:

```go
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
```

Slice && Point:

```
%p
```

## Type alias vs defintion

我们基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias，这就是他们最大的区别。

```
type MyInt1 int
type MyInt2 = int
```

第一行代码是基于基本类型int创建了新类型MyInt1，第二行是创建的一个int的类型别名MyInt2，注意类型别名的定义是`=`。

```
var i int =0
var i1 MyInt1 = i //error
var i2 MyInt2 = i
fmt.Println(i1,i2)
```

仔细看这个示例，第二行把一个`int`类型的变量`i`,赋值给`MyInt1`类型的变量`i1`会被提示编译错误：类型无法转换。但是第三行把`int`类型的变量`i`,赋值给`MyInt2`类型的变量`i2`就可以，不会提示错误。

从这个例子也可以看出来，这两种定义方式的不同，因为Go是强类型语言，所以类型之间的转换必须强制转换，因为`int`和`MyInt1`是不同的类型，所以这里会报编译错误。

但是因为`MyInt2`只是`int`的一个别名，本质上还是一个`int`类型，所以可以直接赋值，不会有问题。

### 类型方法

每个类型都可以通过接受者的方式，添加属于它自己的方法，我们看下通过type alias的类型是否可以，以及拥有哪些方法。

```
type MyInt1 int
type MyInt2 = int
func (i MyInt1) m1(){
    fmt.Println("MyInt1.m1")
}
func (i MyInt2) m2(){
    fmt.Println("MyInt2.m2")
}
func main() {
    var i1 MyInt1
    var i2 MyInt2
    i1.m1()
    i2.m2()
}
```

以上示例代码看着是没有任何问题，但是我们编译的时候会提示：

```
i2.m2 undefined (type int has no field or method m2)
cannot define new methods on non-local type int
```

这里面有2个错误，一个是提示类型`int`没有`m2`这个方法，所以我们不能调用，因为`MyInt2`本质上就是`int`。

第二个错误是我们不能为`int`类型添加新方法，什么意思呢？因为`int`是一个非本地类型，所以我们不能为其增加方法。既然这样，那我们自定义个struct类型试试。

```
type User struct {
}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
    fmt.Println("MyUser1.m1")
}
func (i MyUser2) m2(){
    fmt.Println("MyUser2.m2")
}
//Blog:www.flysnow.org
//Wechat:flysnow_org
func main() {
    var i1 MyUser1
    var i2 MyUser2
    i1.m1()
    i2.m2()
}
```

换成struct，正常运行。所以本地定义的类型的别名，还是可以为其添加方法的。现在我们接着上面的例子，看一个有趣的现象，我在main函数里增加如下代码：

```
var i User
    i.m2()
```

然后运行，发现，可以正常运行。是不是很奇怪，我们并没有为类型`User` 定义方法啊，怎么可以调用呢？这就得益于type alias，`MyUser2`完全等价于`User`，所以为`MyUser2`定义方法，等于就为`User`定义了方法，反之，亦然。

但是对于新定义的类型`MyUser1`就不行了，因为它完全是个新类型，所以`User`的方法，`MyUser`是没有的。这里不再举例，大家自己可以试试。

还有一点需要注意，因为`MyUser2`完全等价于`User`，所以`User`已经有的方法，`MyUser2`不能再声明，反之亦然，如果定义了会有如下提示：

```
./main.go:37:6: User.m1 redeclared in this block
    previous declaration at ./main.go:31:6
```

其实就是重复声明的意思，不能再次重复声明了。

### 接口实现

上面的小结我们可以发现，`User`和`MyUser2`是等价的，并且其中一个新增了方法，另外一个也会有。那么基于此推导出，一个实现了某个接口，另外一个也会实现。现在验证一下：

```
type I interface {
    m2()
}
type User struct {
}
type MyUser2 = User
func (i User) m(){
    fmt.Println("User.m")
}
func (i MyUser2) m2(){
    fmt.Println("MyUser2.m2")
}
func main() {
    var u User
    var u2 MyUser2
    var i1 I =u
    var i2 I =u2
    fmt.Println(i1,i2)
}
```

定义了一个接口`I`，从代码上看，只有`MyUser2`实现了它，但是我们代码的演示中，发现`User`也实现了接口`I`，所以这就验证了我们的推到是正确的，返回来如果`User`实现了某个接口，那么它的type alias也同样会实现这个接口。

以上讲了很多示例都是类型struct的别名，我们看下接口interface的type alias是否也是等价的。

```
type I interface {
    m()
}
type MyI1 I
type MyI2 = I
type MyInt int
func (i MyInt) m(){
    fmt.Println("MyInt.m")
}
```

定义了一个接口`I`，`MyI1`是基于`I`的新类型；`MyI2`是`I`的类型别名；`MyInt`实现了接口`I`。下面进行测试。

```
func main() {
    //赋值实现类型MyInt
    var i I = MyInt(0)
    var i1 MyI1 = MyInt(0)
    var i2 MyI2 = MyInt(0)
    //接口间相互赋值
    i = i1
    i = i2
    i1 = i2
    i1 = i
    i2 = i
    i2 = i1
}
```

以上代码运行是正常的，这个是前面讲的具体类型（struct，int等）的type alias不一样，只要实现了接口，就可以相互赋值，管你是新定义的接口`MyI1`，还是接口的别名`MyI2`。

### 类型的嵌套

我们都知道type alias的两个类型是等价的，但是他们在类型嵌套时有些不一样。

```
func main() {
    my:=MyStruct{}
    my.T2.m1()
}
type T1 struct {
}
func (t T1) m1(){
    fmt.Println("T1.m1")
}
type T2 = T1
type MyStruct struct {
    T2
}
```

示例中`T2`是`T1`的别名，但是我们把T2嵌套在`MyStruct`中，在调用的时候只能通过`T2`这个名称调用，而不能通过`T1`，会提示没这个字段的。反过来也一样。

这是因为`T1`,`T2`是两个名称，虽然他们等价，但他们是具有两个不同名字的等价类型，所以在类型嵌套的时候，就是两个字段。

当然我们可以把`T1`,`T2`同时嵌入到`MyStrut`中，进行分别调用。

```
func main() {
    my:=MyStruct{}
    my.T2.m1()
    my.T1.m1()
}
type MyStruct struct {
    T2
    T1
}
```

以上也是可以正常运行的，证明这是具有两个不同名字的，同种类型的字段。

下面我们做个有趣的实验，把`main`方法的代码改为如下：

```
func main() {
    my:=MyStruct{}
    my.m1()
}
```

猜猜是不是可以正常编译运行呢？答应可能出乎意料，是不能正常编译的，提示如下：

```
./main.go:25:4: ambiguous selector my.m1
```

其实想想很简单，不知道该调用哪个，太模糊了，匹配不了，不然该用`T1`的`m1`,还是`T2`的`m1`。这种结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。

### 类型循环

type alias的声明，一定要留意类型循环，不要产生了循环，一旦产生，就会编译不通过，那么什么是类型循环呢。假如`type T2 = T1`,那么`T1`绝对不能直接、或者间接的引用到`T2`，一旦有，就会类型循环。

```
type T2 = *T2
type T2 = MyStruct
type MyStruct struct {
    T1
    T2
}
```

以上两种定义都是类型循环，我们自己在使用的过程中，要避免这种定义的出现。

byte and rune

这两个类型一个是int8的别名，一个是int32的别名，在Go 1.9之前，他们是这么定义的。

```
type byte byte

type rune rune
```

现在Go 1.9有了type alias这个新特性后，他们的定义就变成如下了：

```
type byte = uint8

type rune = int32
```

恩，非常很省事和简洁。

### 导出未导出的类型

type alias还有一个功能，可以导出一个未被导出的类型。

```
package lib

type user struct {
    name string
    Email string
}
func (u user) getName() string {
    return u.name
}
func (u user) GetEmail() string {
    return u.Email
}
//把这个user导出为User
type User = user
```

`user`本身是一个未导出的类型，不能被其他package访问，但是我们可以通过`type User = user`，定义一个`User`，这样这个`User`就可以被其他package访问了，可以使用`user`类型导出的字段和方法，示例中是`Email`字段和`GetEmail`方法，另外未被导出`name`字段和`getName`方法是不能被其他package使用的。

小结

type alias的定义，本质上是一样的类型，只是起了一个别名，源类型怎么用，别名类型也怎么用，保留源类型的所有方法、字段等。

## Go语言基础语法

### 包

每个 Go 程序都是由包构成的。

程序从 `main` 包开始运行。

本程序通过导入路径 `"fmt"` 和 `"math/rand"` 来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。例如，`"math/rand"` 包中的源码均以 `package rand` 语句开始。

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

### 导入

此代码用圆括号组合了导入，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

```
import "fmt"
import "math"
```

不过使用分组导入语句是更好的形式。

```
import (
	"fmt"
	"math"
)
```

导入相对路径的package 报错：

```
go build
main.go:6:2: local import "./testinit" in non-local package
```

如果把project移出gopath，这个时候就不会报这个编译错误了。
**golang只有在gopath中找不到的包路径，才允许用相对路径导入**。

### 特殊用法

1 使用点操作引入包时，可以省略包前缀：

2 别名 import f "fmt"

3 _操作

```go
// 对包 lib 的调用直接省略包名
import . "utils/lib"

//  别名
import f "fmt"

// 特殊符号“_” 仅仅会导致 lib 执行初始化工作，如初始化全局变量，调用 init 函数。
import _ "utils/lib"
```



### package 的init 方法

引入一个package的时候，如果有init方法，会自动执行，用作初始化。

1. 在同一个package中，可以多个文件中定义init方法
2. 在同一个go文件中，可以重复定义init方法
3. 在同一个package中，不同文件中的init方法的执行按照文件名先后执行各个文件中的init方法
4. 在同一个文件中的多个init方法，按照在代码中编写的顺序依次执行不同的init方法

### 导出名

- 大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
- 大写字母开头的函数也是一样，相当于`class`中的带`public`关键词的公有函数；小写字母开头的就是有`private`关键词的私有函数。

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。 例如， `Pizza` 就是个已导出名， `Pi` 也同样，它导出自 `math` 包。

`pizza` 和 `pi` 并未以大写字母开头，所以它们是未导出的。

在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

执行代码，观察错误输出。

然后将 `math.pi` 改名为 `math.Pi` 再试着执行一次。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```

### nil

nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.

```go
 // Type must be a pointer, channel, func, interface, map, or slice type
```

在 golang 中，nil 只能赋值给 指针、channel、func、interface、map 或 slice 类型的变量。如果未遵循这个规则，则会引发 panic

### 内置函数

Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作，例如：len、cap 和 append，或必须用于系统级的操作，例如：panic。因此，它们需要直接获得编译器的支持。

以下是一个简单的列表，我们会在后面的章节中对它们进行逐个深入的讲解。

| 名称               | 说明                                                         |
| ------------------ | ------------------------------------------------------------ |
| close              | 用于管道通信                                                 |
| len、cap           | len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map） |
| new、make          | new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：`v := new(int)`。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作。**new() 是一个函数，不要忘记它的括号** |
| copy、append       | 用于复制和连接切片                                           |
| panic、recover     | 两者均用于错误处理机制                                       |
| print、println     | 底层打印函数，在部署环境中建议使用 fmt 包                    |
| complex、real imag | 用于创建和操作复数                                           |

### 函数

Go 语言函数定义格式如下：

```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

函数可以没有参数或接受多个参数。

在本例中， `add` 接受两个 `int` 类型的参数。

注意类型在变量名 **之后** 。

```go
func add(x int, y int) int {
	return x + y
}
```

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略

```go
x int, y int 
// 缩写为
x, y int
```

### 函数多值返回

函数可以返回任意数量的返回值。

`swap` 函数返回了两个字符串。

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

### 函数命名返回值

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 `return` 语句返回已命名的返回值。也就是`直接`返回。

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

### 函数参数传递

调用函数，可以通过两种方式来传递参数：

| 传递类型                                                     | 描述                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [值传递](http://www.runoob.com/go/go-function-call-by-value.html) | 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。 |
| [引用传递](http://www.runoob.com/go/go-function-call-by-reference.html) | 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。 |

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

**Go语言中所有的传参都是值传递（传值）**，都是一个副本，一个拷贝。因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样就在函数中就无法修改原内容数据；**有的是引用类型（指针、map、slice、chan等这些），这样就可以修改原内容数据。**

是否可以修改原内容数据，和传值、传引用没有必然的关系。在C++中，传引用肯定是可以修改原内容数据的，在Go语言里，**虽然只有传值，但是我们也可以修改原内容数据，因为参数是引用类型。**

引用类型和传引用是两个概念。

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )

   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)

   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```



### 变量

`var` 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，`var` 语句可以出现在包或函数级别。

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

```



### 变量的初始化

`var` 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

`var` 语句可以出现在包或函数级别。

变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

```go
package main

import "fmt"

var i, j int = 1, 2
var arr = [2]int{1,2}
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
  fmt.Println(arr)
}

```

### 短变量声明

在函数中，简洁赋值语句 `:=` 可在类型明确的地方代替 `var` 声明。

函数外的每个语句都必须以关键字开始（`var`, `func` 等等），因此 `:=` 结构不能在函数外使用。

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java) 
}
```

`_`（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。

将值`35`赋予`b`，并同时丢弃`34`：

```go
_, b := 34, 35
```

### 变量作用域

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

### 内置基础类型

#### Boolean

在Go中，布尔值的类型为`bool`，值是`true`或`false`，默认为`false`。

```Go
//示例代码
var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明
func test() {
    var available bool  // 一般声明
    valid := false      // 简短声明
    available = true    // 赋值操作
}
```

#### 整数类型

整数类型有无符号和带符号两种。Go同时支持`int`和`uint`，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：`rune`, `int8`, `int16`, `int32`, `int64`和`byte`, `uint8`, `uint16`, `uint32`, `uint64`。其中`rune`是`int32`的别称，`byte`是`uint8`的别称。


byte`用来强调数据是`raw data`，而不是数字；而`rune`用来表示`Unicode`的`code point


这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错

```go
var a int8

var b int32

c:=a + b

//会产生错误：invalid operation: a + b (mismatched types int8 and int32)
```

#### 字符串

Go中的字符串都是采用`UTF-8`字符集编码。字符串是用一对双引号（`""`）或反引号（`` ` `` `）括起来定义，它的类型是`string`。

在Go中字符串是不可变的，例如下面的代码编译时会报错：cannot assign to s[0]

```Go
var s string = "hello"
s[0] = 'c'
```

但如果真的想要修改怎么办呢？下面的代码可以实现：

```Go
s := "hello"
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```

Go中可以使用`+`操作符来连接两个字符串：

```Go
s := "hello,"
m := " world"
a := s + m
fmt.Printf("%s\n", a)
```

修改字符串也可写为：

```Go
s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)
```

如果要声明一个多行的字符串怎么办？可以通过`` `来声明：

```
m := `hello
    world`
```

`` ` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。例如本例中会输出：

```
hello
    world
```

#### 错误类型

Go内置有一个`error`类型接口，专门用来处理错误信息。

```go
type error interface {
        Error() string
}
```

Go的`package`里面还专门有一个包`errors`来处理错误：

```Go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
    fmt.Print(err)
}
```

#### 类型转换

表达式 `T(v)` 将值 `v` 转换为类型 `T` 。

```go
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
<-chan int(c)    // same as <-(chan int(c))
(<-chan int)(c)  // c is converted to <-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
```

一些关于数值的转换：

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

或者，更加简单的形式：

```go
i := 42
f := float64(i)
u := uint(f)
```

如果对于某些地方的优先级拿不准可以自己加`()`约束.

与 C 不同的是，**Go 在不同类型的项之间赋值时需要显式转换**

#### 类型推导

在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法 或 `var =` 表达式语法）， 变量的类型由右值推导得出。

当右值声明了类型时，新变量的类型与其相同：

```
var i int
j := i // j 也是一个 int
```

不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 `int` 、 `float64`或 `complex128` 了，这取决于常量的精度：

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

尝试修改示例代码中 `v` 的初始值，并观察它是如何影响类型的。

```go
package main

import "fmt"

func main() {
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
}
```

#### 常量

所谓常量，也就是在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。

常量的声明与变量类似，只不过是使用 `const` 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 `:=` 语法声明

```go
const constantName = value
//如果需要，也可以明确指定常量的类型：
const Pi float32 = 3.1415926
```

#### iota枚举

Go里面有一个关键字`iota`，这个关键字用来声明`enum`的时候采用，它默认开始值是0，const中每增加一行加1：

```Go
const (
    x = iota // x == 0
    y = iota // y == 1
    z = iota // z == 2
    w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
    h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
    a       = iota //a=0
    b       = "B"
    c       = iota             //c=2
    d, e, f = iota, iota, iota //d=3,e=3,f=3
    g       = iota             //g = 4
)
```

#### 错误类型

Go内置有一个`error`类型，专门用来处理错误信息，Go的`package`里面还专门有一个包`errors`来处理错误：

```Go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
    fmt.Print(err)
}
```

#### Type

```go
type test_int func(int) bool //申明了一个函数类型
```



### 基础流程语法

#### 分组声明

可以分组写成如下形式：

```Go
import(
    "fmt"
    "os"
)

const(
    i = 100
    pi = 3.1415
    prefix = "Go_"
)

var(
    i int
    pi float32
    prefix string
)
```

#### for

Go 只有一种循环结构： `for` 循环。

基本的 `for` 循环由三部分组成，它们用分号隔开：

- 初始化语句：在第一次迭代前执行
- 条件表达式：在每次迭代前求值
- 后置语句：在每次迭代的结尾执行

这三部分组成的循环的头部，它们之间使用分号 `;` 相隔，但并不需要括号 `()` 将它们括起来

初始化语句通常为一句短变量声明，该变量声明仅在 `for` 语句的作用域中可见。

一旦条件表达式的布尔值为 `false`，循环迭代就会终止。 **注意**：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面没有小括号，大括号 `{ }` 则是必须的。

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

左花括号 `{` 必须和 for 语句在同一行，计数器的生命周期在遇到右花括号 `}` 时便终止。一般习惯使用 i、j、z 或 ix 等较短的名称命名计数器。

特别注意，永远不要在循环体内修改计数器，这在任何语言中都是非常差的实践！





```go
package main

import "fmt"

func main() {
    str := "Golang"
    fmt.Printf("The length of str is: %d\n", len(str))
    for ix :=0; ix < len(str); ix++ {
        fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
    }
    str2 := "日本語"
    fmt.Printf("The length of str2 is: %d\n", len(str2))
    for ix :=0; ix < len(str2); ix++ {
        fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
    }
}
```



#### for 是 Go 中的 “while”

此时你可以去掉分号，因为 C 的 `while` 在 Go 中叫做 `for` 。

```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

#### for-range 结构

这是 Go 特有的一种的迭代结构，您会发现它在许多情况下都非常有用。它可以迭代任何一个集合（包括数组和 map，详见第 7 和 8 章）。语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。一般形式为：`for ix, val := range coll { }`。

要注意的是，`val` 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值（**译者注：如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值**）。一个字符串是 Unicode 编码的字符（或称之为 `rune`）集合，因此您也可以用它迭代字符串：

```go
package main

import "fmt"

func main() {
    str := "Golang!"
    fmt.Printf("The length of str is: %d\n", len(str))
    for pos, char := range str {
        fmt.Printf("Character on position %d is: %c \n", pos, char)
    }
    fmt.Println()
    str2 := "Chinese: 日本語"
    fmt.Printf("The length of str2 is: %d\n", len(str2))
    for pos, char := range str2 {
        fmt.Printf("character %c starts at byte position %d\n", char, pos)
    }
    fmt.Println()
    fmt.Println("index int(rune) rune    char bytes")
    for index, rune := range str2 {
        fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
    }
}
```

每个 rune 字符和索引在 for-range 循环中是一一对应的。它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。

#### if

Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 `( )` ，而大括号 `{ }` 则是必须的。

```go
if x < 0 {
	return sqrt(-x) + "i"
}
```

#### if 的简短语句

同 `for` 一样， `if` 语句可以在条件表达式前执行一个简单的语句。

该语句声明的变量作用域仅在 `if` 之内。

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

#### if 和 else

在 `if` 的简短语句中声明的变量同样可以在任何对应的 `else` 块中使用。

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}
```

#### switch

- `switch` 关键字后面接条件表达式
- `case` 从上到下按顺序进行匹配，直到匹配成功
- 如果没有匹配到 `case`, 且有 `default` 模式， 会执行 `default` 的代码块

```go
func defaultSwitch() {
    switch time.Now().Weekday() {
    case time.Saturday:
        fmt.Println("Today is Saturday.")
    case time.Sunday:
        fmt.Println("Today is Sunday.")
    default:
        fmt.Println("Today is a weekday.")
    }
}

func swi(){
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
}

```

你大概已经知道 `switch` 语句的样子了。

```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。前花括号 `{` 必须和 switch 关键字在同一行。

您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：`case val1, val2, val3`。

每一个 `case` 分支都是唯一的，从上至下逐一测试，直到匹配为止。

一旦成功地匹配到每个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 `break` 语句来表示结束。

因此，程序也不会自动地去执行下一个分支的代码。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 `fallthrough` 关键字来达到目的。

golang中的fallthrough用在switch的case中，case执行完之后一般break，但可以使用 fallthrough 来强制执行下一个 case 代码块

o里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码(就算后面的case 是false)。

```
    i := 2
    switch i {
    case 1:
        fmt.Println("one")
        fallthrough
    case 2:
        fmt.Println("two")
        fallthrough
    case 3:
        fmt.Println("three")
    }
    
    
    two
    three // i!=3
    
```



```go
package main

import "fmt"


func main() {
    switch {
    case false:
        fmt.Println("The integer was <= 4")
        fallthrough
    case true:
        fmt.Println("The integer was <= 5")
        fallthrough
    case false:
        fmt.Println("The integer was <= 6")
        fallthrough
    case true:
        fmt.Println("The integer was <= 7")
        fallthrough
    case false:
        fmt.Println("The integer was <= 8")
    default:
        fmt.Println("default case")
    }
}

The integer was <= 5
The integer was <= 6
The integer was <= 7
The integer was <= 8
```



```go
switch i {
    case 0: // 空分支，只有当 i == 0 时才会进入分支
    case 1:
        f() // 当 i == 0 时函数不会被调用
}
```

```go
switch i {
    case 0: fallthrough
    case 1:
        f() // 当 i == 0 时函数也会被调用
}
```

除非以 `fallthrough` 语句结束，否则分支会自动终止。fallthrough不能用在switch的最后一个分支。

```go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
```

##### switch 的求值顺序

switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。



#### 没有条件的 switch

没有条件的 switch 同 `switch true` 一样。

这种形式能将一长串 if-then-else 写得更加清晰。

```go
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

```go
package main

import "fmt"

func main() {
    var num1 int = 7

    switch {
        case num1 < 0:
            fmt.Println("Number is negative")
        case num1 >= 0 && num1 < 10:
            fmt.Println("Number is between 0 and 10")
        default:
            fmt.Println("Number is 10 or greater")
    }
  
   executeOrder()
}

func executeOrder() {
    // Foo prints and returns n.
    var Foo func(int) int
    Foo = func(n int) int {
        fmt.Println(n)
        return n
    }

    switch Foo(2) {
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
       // fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}

// 2
// 1 2
```



#### defer

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
//hello
//world
```

#### defer 栈

推迟的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

更多关于 defer 语句的信息， 请阅读[此博文](http://blog.go-zh.org/defer-panic-and-recover)。

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
/*
结果：
counting
done
9
8
7
6
5
4
3
2
1
0
*/

```

## 运算符

#### 算术运算符

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例               |
| :----- | :--- | :----------------- |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |

#### 关系运算符

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |



#### 逻辑运算符

下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。

| 运算符 | 描述                                                         | 实例               |
| :----- | :----------------------------------------------------------- | :----------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A \|\| B) 为 True |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 | !(A && B) 为 True  |



#### 位运算符

位运算符对整数在内存中的二进制位进行操作。

Go 语言提供了以下的 bit 位操作运算符

```
& 位运算 AND

| 位运算 OR

^ 位运算 XOR

&^ 位清空 (AND NOT)
《Go语言圣经》里面有对此的描述，z = x &^ y，如果 y 中比特位为 1 那么 z 对应比特位值就是 0，否则 z 就使用 x 对应位置的比特位值。
实际执行方式是 x 和 (y取反后的值)进行与操作

<< 左移

<< 左移
```







```
const (
		Read byte = 1 << iota
		Write
		Execute
	)

	var f1 = Read | Write | Execute
	var f2 = Read | Write

	var f = f1 &^ f2

	fmt.Printf("%03b &^ %03b = %03b\n", f1, f2, f)
	
	// 111 &^ 011 = 100
```

Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

| 运算符 | 描述                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | 按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 | A >> 2 结果为 15 ，二进制为 0000 1111  |



#### 赋值运算符

下表列出了所有Go语言的赋值运算符。

| 运算符 | 描述                                           | 实例                                  |
| :----- | :--------------------------------------------- | :------------------------------------ |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C |
| +=     | 相加后再赋值                                   | C += A 等于 C = C + A                 |
| -=     | 相减后再赋值                                   | C -= A 等于 C = C - A                 |
| *=     | 相乘后再赋值                                   | C *= A 等于 C = C * A                 |
| /=     | 相除后再赋值                                   | C /= A 等于 C = C / A                 |
| %=     | 求余后再赋值                                   | C %= A 等于 C = C % A                 |
| <<=    | 左移后赋值                                     | C <<= 2 等于 C = C << 2               |
| >>=    | 右移后赋值                                     | C >>= 2 等于 C = C >> 2               |
| &=     | 按位与后赋值                                   | C &= 2 等于 C = C & 2                 |
| ^=     | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                 |
| \|=    | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2               |

#### golang逗号模式

##### 在函数返回时检测错误

```go
// os.Open(file) strconv.Atoi(str)
value, err := pack1.Func1(param1)

if err != nil {
    fmt.Printf(“Error %s in pack1.Func1 with parameter %v”, err.Error(), param1)
    return err
}
```

##### 检测映射中是否存在一个键值5

```go
if value, isPresent = map1[key1]; isPresent {
        Process(value)
}
// key1不存在
…
```

##### 检测一个接口类型变量varI是否包含了类型T

```go
if value, ok := varI.(T); ok {
    Process(value)
}
// 接口类型varI没有包含类型T
```

##### 检测一个通道ch是否关闭

```go
 for {
        if input, open := <-ch; !open {
            break // 通道是关闭的
        }
        Process(input)
    }
```





## 结构体struct

一个结构体（`struct`）就是一个字段的集合。

```go
type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(Vertex{1, 2})
}
```

### 结构体字段

结构体字段使用点号来访问。

### 结构体指针

结构体字段可以通过结构体指针来访问。

如果我们有一个指向结构体的指针 `p` ，那么可以通过 `(*p).X` 来访问其字段 `X` 。 不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 `p.X` 就可以。

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

Go支持只提供类型，而不写字段名的方式，也就是**匿名字段**，也称为嵌入字段。

当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。

```go
type Human struct {
    name string
    age int
    weight int
}

type Student struct {
    Human  // 匿名字段，那么默认Student就包含了Human的所有字段
    speciality string
}
```

通过匿名访问和修改字段相当的有用，但是不仅仅是struct字段，所有的内置类型和自定义类型都是可以作为匿名字段的

```go
type Student struct {
    Human  // 匿名字段，那么默认Student就包含了Human的所有字段
    speciality string
    int    // 内置类型作为匿名字段
}
```

```go
package main

import "fmt"

type Human struct {
    name string
    age int
    weight int
}
type Student struct {
    Human  // 匿名字段，struct
    int    // 内置类型作为匿名字段
    speciality string
    int32 // 不能是int 重复 
}
func main() {
    jane := Student{Human:Human{"Jane", 35, 100}, speciality:"Biology"}
    jane.int = 15
    fmt.Println(jane)
}
// {{Jane 35 100} 15 Biology 0}
```



### struct成员变量标签tag

在处理json格式字符串的时候，经常会看到声明struct结构的时候，属性的右侧还有小米点括起来的内容。形如

```go
type User struct {
    UserId   int    `json:"user_id" bson:"user_id"`
    UserName string `json:"user_name" bson:"user_name"`
}
```

在golang中，命名都是推荐都是用驼峰方式，并且在首字母大小写有特殊的语法含义：包外无法引用。但是由经常需要和其它的系统进行数据交互，例如转成json格式，存储到mongodb啊等等。这个时候如果用属性名来作为键值可能不一定会符合项目要求。

所以呢就多了小米点的内容，在golang中叫标签（Tag），在转换成其它数据格式的时候，会使用其中特定的字段作为键值。例如上例在转成json格式：

```go
u := &User{UserId: 1, UserName: "tony"}
j, _ := json.Marshal(u)
fmt.Println(string(j))
// 输出内容：{"user_id":1,"user_name":"tony"}

// 如果在属性中不增加标签说明，则输出：
// {"UserId":1,"UserName":"tony"}
```

可以看到直接用struct的属性名做键值。

### struct成员变量标签Tag获取

用反射包（reflect）中的方法来获取

```go
package main
 
import (
    "encoding/json"
    "fmt"
    "reflect"
)
 
func main() {
    type User struct {
        UserId   int    `json:"user_id" bson:"user_id"`
        UserName string `json:"user_name" bson:"user_name"`
    }
    // 输出json格式
    u := &User{UserId: 1, UserName: "tony"}
    j, _ := json.Marshal(u)
    fmt.Println(string(j))
    // 输出内容：{"user_id":1,"user_name":"tony"}
 
    // 获取tag中的内容
    t := reflect.TypeOf(u)
    elem := t.Elem();
    field := elem.Field(0)
    fmt.Println(field.Tag.Get("json"))
    // 输出：user_id
    fmt.Println(field.Tag.Get("bson"))
    // 输出：user_id
    
    for i := 0; i < elem.NumField(); i++ {
        fmt.Println(elem.Field(i).Tag) //将tag输出出来
    }
}
```



## 数组

### 数组

类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

```Go
var arr [n]type
```

在`[n]type`中，`n`表示数组的长度，`type`表示存储元素的类型。对数组的操作和其它语言类似，都是通过`[]`来进行读取或赋值：

表达式

```
var a [10]int
```

会将变量 `a` 声明为拥有有 10 个整数的数组。

由于长度也是数组类型的一部分，因此数组不能改变大小，例如`[3]int`与`[4]int`是不同的类型，数组也就不能改变长度。数组之间的赋值是值的赋值

当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么就需要用到后面介绍的`slice`类型了

```Go
var a = [3]int{1, 2, 3} // 声明了一个长度为3的int数组

a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

d:= [...]int{1:1,0:2}
```
二维数组：

```Go
// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

// 上面的声明可以简化，直接忽略内部的类型
easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
```



```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
    iarray3 := [5]int32{1, 2, 3, 4, 5}
    iarray4 := []int32{1, 2, 3, 4, 5}
    iarray5 := [...]int32{1, 2, 3, 4, 5}
    fmt.Printf("%T,%T,%T", iarray3, iarray4 ,iarray5)
}
```

### 切片

在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫`slice`

`slice`并不是真正意义上的动态数组，而是一个引用类型。`slice`总是指向一个底层`array`，`slice`的声明也可以像`array`一样，只是不需要长度。

每个数组的大小都是固定的。 而切片则为数组元素提供动态大小的、灵活的视角。 在实践中，切片比数组更常用。

切片的创建有**4**种方式：

1）make ( []Type ,length, capacity )

2)  make ( []Type, length)

3) []Type{}

4) []Type{value1 , value2 , ... , valueN }

类型 `[]T` 表示一个元素类型为 `T` 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```
a[low : high]
```

它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```go
a[1:4]
```

```go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

`slice`可以从一个数组或一个已经存在的`slice`中再次声明。`slice`通过`array[i:j]`来获取，其中`i`是数组的开始位置，`j`是结束位置，但不包含`array[j]`，它的长度是`j-i`。

下面这个例子展示了更多关于`slice`的操作：

```Go
// 声明一个数组
var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
// 声明两个slice
var aSlice, bSlice []byte

// 演示一些简便操作
aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

// 从slice中获取slice
aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
```

### 切片就像数组的引用

切片并不存储任何数据， 它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
/*
result:
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
*/

```

### 切片文法

切片文法类似于没有长度的数组文法。

这是一个数组文法：

```
[3]bool{true, true, false}
```

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

```
[]bool{true, true, false}
```

```
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```
### 切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 `0` ，上界则是该切片的长度。

对于数组

```
var a [10]int
```

来说，以下切片是等价的：

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

### 切片的长度与容量

切片拥有 **长度** 和 **容量** 。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

你可以通过重新切片来扩展一个切片。

### nil 切片

切片的零值是 `nil` 。

nil 切片的长度和容量为 0 且没有底层数组。

### 用 make 创建切片

切面可以用内建函数 `make` 来创建，这也是你创建动态数组的方式。

`make` 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

```
a := make([]int, 5)  // len(a)=5
```

要指定它的容量，需向 `make` 传入第三个参数：

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	
	e := b[:cap(b)]
	printSlice("e", e)
	
	f := e[1:]
	printSlice("f", f)
	
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
e len=5 cap=5 [0 0 0 0 0]
f len=4 cap=4 [0 0 0 0]
*/

```

### 切片的切片

切片可包含任何类型，甚至包括其它的切片。

### 向切片追加元素

为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 `append` 函数。 内建函数的[文档](https://go-zh.org/pkg/builtin/#append)对此函数有详细的介绍。

```
func append(s []T, vs ...T) []T
```

`append` 的第一个参数 `s` 是一个元素类型为 `T` 的切片， 其余类型为 `T` 的值将会追加到该切片的末尾。

`append` 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 `s` 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。 返回的切片会指向这个新分配的数组。

（要了解关于切片的更多内容，请阅读文章[Go 切片：用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals)。）

### Range

`for` 循环的 `range` 形式可遍历切片或映射。

当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

```go
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

可以将下标或值赋予 `_` 来忽略它。

若你只需要索引，去掉 `, value` 的部分即可。

```go
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

### 通过切片模拟栈和队列



栈

```
// 创建栈
stack:=make([]int,0)
// push压入
stack=append(stack,10)
// pop弹出
v:=stack[len(stack)-1]
stack=stack[:len(stack)-1]
// 检查栈空
len(stack)==0
```

队列

```
// 创建队列
queue:=make([]int,0)
// enqueue入队
queue=append(queue,10)
// dequeue出队
v:=queue[0]
queue=queue[1:]
// 长度0为空
len(queue)==0
```

注意点

- 参数传递，只能修改，不能新增或者删除原始数据
- 默认 s=s[0:len(s)]，取下限不取上限，数学表示为：[)

## 映射(Map)

映射的创建的4种方式

make ( map [KeyType] ValueType, initialCapacity )

make ( map [KeyType] ValueType )

map [KeyType ] ValueType {}

map [KeyType ] ValueType { key1 : value1, key2: value2, ... , keyN : valueN}

映射将键映射到值。

映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。

```go
// 这种方式初始化的不能直接用
var m1 map[string]string
// panic: assignment to entry in nil map
m1["key"] = "value"

// 正确用法
// 先声明map
var m1 map[string]string
// 再使用make函数创建一个非nil的map，nil map不能赋值
m1 = make(map[string]string)
// 最后给已声明的map赋值
m1["a"] = "aa"

// 直接创建
m2 := make(map[string]string)
// 然后赋值
m2["a"] = "aa"
m2["b"] = "bb"

// 初始化 + 赋值一体化
m3 := map[string]string{
	"a": "aa",
	"b": "bb",
}
```

`make` 函数会返回给定类型的映射，并将其初始化备用。

`map`的读取和设置也类似`slice`一样，通过`key`来操作，只是`slice`的`index`只能是｀int｀类型，而`map`多了很多类型，可以是`int`，可以是`string`及所有完全定义了`==`与`!=`操作的类型。

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
    var mn = map[string]Vertex{
    	"Bell Labs": Vertex{
    		40.68433, -74.39967,
    	},
    	"Google": Vertex{
    		37.42202, -122.08408,
    	},
    }
    fmt.Println(mn)
}
```

## 修改映射

在映射 `m` 中插入或修改元素：

```
m[key] = elem
```

获取元素：

```
elem = m[key]
```

删除元素：

```
delete(m, key)
```

通过双赋值检测某个键是否存在：

```
elem, ok = m[key]
```

若 `key` 在 `m` 中， `ok` 为 `true` ；否则， `ok` 为 `false` 。

若 `key` 不在映射中，那么 `elem` 是该映射元素类型的零值。

同样的，当从 映射 中读取某个不存在的键时，结果是 映射 的元素类型的零值。

**注** ：若 `elem` 或 `ok` 还未声明，你可以使用短变量声明：

```
elem, ok := m[key]
```

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

使用map过程中需要注意的几点：

- `map`是无序的，每次打印出来的`map`都会不一样，它不能通过`index`获取，而必须通过`key`获取
- `map`的长度是不固定的，也就是和`slice`一样，也是一种引用类型
- 内置的`len`函数同样适用于`map`，返回`map`拥有的`key`的数量
- `map`的值可以很方便的修改，通过`numbers["one"]=11`可以很容易的把key为`one`的字典值改为`11`
- `map`和其他基本型别不同，它不是thread-safe，**在多个go-routine存取时，必须使用mutex lock机制**



#### make、new操作

`make`用于内建类型（`map`、`slice` 和`channel`）的内存分配。`new`用于各种类型的内存分配。

内建函数`new`本质上说跟其它语言中的同名函数功能一样：`new(T)`分配了零值填充的`T`类型的内存空间，并且返回其地址，即一个`*T`类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型`T`的零值。有一点非常重要：

> `new`返回指针。

内建函数`make(T, args)`与`new(T)`有着不同的功能，make只能创建`slice`、`map`和`channel`，并且返回一个有初始值(非零)的`T`类型，而不是`*T`。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个`slice`，是一个包含指向数据（内部`array`）的指针、长度和容量的三项描述符；在这些项目被初始化之前，`slice`为`nil`。对于`slice`、`map`和`channel`来说，`make`初始化了内部的数据结构，填充适当的值。

> `make`返回初始化后的（非零）值。



## 函数值

函数也是值。它们可以像其它值一样传递。

函数值可以用作函数的参数或返回值。

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

```

## 函数的闭包

Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。 该函数可以访问并赋予其引用的变量的值，换句话说，该函数被“绑定”在了这些变量上。

例如，函数 `adder` 返回一个闭包。每个闭包都被绑定在其各自的 `sum` 变量上。

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

## 方法

方法就是一个包含了接受者的函数

Go 没有类。不过你可以为结构体类型定义方法，也可以为非结构体类型声明方法(见下面解释)

```go
func (r ReceiverType) funcName(parameters) (results)
```

方法就是一类带特殊的 **接收者** 参数的函数。

方法接收者在它自己的参数列表内，位于 `func` 关键字和方法名之间。

在此例中， `Abs` 方法拥有一个名为 `v` ，类型为 `Vertex` 的接收者。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

## 方法即函数

记住：方法只是个带接收者参数的函数。

现在这个 `Abs` 的写法就是个正常的函数，功能并没有什么变化。

```go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

## 方法（续）

你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 `Abs` 方法的数值类型 `MyFloat` 。

你只能为在同一包内定义的类型的接收者声明方法， 而不能为其它包内定义的类型（包括 `int` 之类的内建类型）的接收者声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）

```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```



## 指针

Go 拥有指针。 指针保存了值的内存地址。

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil` 。

```
var p *int
```

`&` 操作符会生成一个指向其操作数的指针。

```
i := 42
p = &i
```

`*` 操作符表示指针指向的底层值。

```
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

这也就是通常所说的“间接引用”或“重定向”。

与 C 不同，Go 没有指针运算。

```go
package main

import "fmt"

func main() {
    
    
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
/*
result:
42
21
73
*/
```

- 在 `Go` 中 `*` 代表取指针地址中存的值，`&` 代表取一个值的地址

- 对于指针，我们一定要明白指针储存的是一个值的地址，但本身这个指针也需要地址来储存

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	var p *int
  	p = new(int)
  	*p = 1
  	fmt.Println(p, &p, *p)
  }
  
  输出
  0xc04204a080  0xc042068018  1
  
  ```

  - 如上 `p` 是一个指针，他的值为[内存地址](https://www.baidu.com/s?wd=%E5%86%85%E5%AD%98%E5%9C%B0%E5%9D%80&tn=24004469_oem_dg&rsv_dl=gh_pl_sl_csd) `0xc04204a080`
  - 而 `p` 的内存地址为 `0xc042068018`
  - 内存地址 `0xc04204a080` 储存的值为 `1`

**错误实例**

在 **golang** 中如果我们定义一个指针并像普通变量那样给他赋值，例如下方的代码

```go
package main

import "fmt"
func main() {
	var i *int
	*i = 1
    fmt.Println(i, &i, *i)
}

```

- 就会报这样的一个错误

  ```go
  panic: runtime error: invalid memory address or nil pointer dereference
  [signal 0xc0000005 code=0x1 addr=0x0 pc=0x498025]
  ```

- 这个错的原因是 `go` 初始化指针的时候会为指针 `i` 的值赋为 `nil` ，但 `i` 的值代表的是 `*i`的地址， `nil` 的话系统还并没有给 `*i` 分配地址，所以这时给 `*i` 赋值肯定会出错

- 解决这个问题非常简单，在给指针赋值前可以先**创建一块内存**分配给赋值对象即可

  `i = new(int)`

  

## 指针接收者

你可以为指针接收者声明方法。

这意味着对于某类型 `T` ，接收者的类型可以用 `*T` 的文法。 （此外， `T` 不能是像 `*int` 这样的指针。）

例如，这里为 `*Vertex` 定义了 `Scale` 方法。

指针接收者的方法可以修改接收者指向的值（就像 `Scale` 在这做的）。 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

试着移除第 16 行 `Scale` 函数声明中的 `*` ，观察此程序的行为如何变化。

若使用值接收者，那么 `Scale` 方法会对原始 `Vertex` 值的副本进行操作。 （对于函数的其它参数也是如此。） `Scale` 方法必须用指针接受者来更改 `main` 函数中声明的 `Vertex` 的值。

方法的Receiver是以值传递，而非引用传递，是的，Receiver还可以是指针, 两者的差别在于, 指针作为Receiver会对实例对象的内容发生操作,而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

## 指针与函数

现在我们要把 `Abs` 和 `Scale` 方法重写为函数。

```go
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```

## 方法与指针重定向

比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

```
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK
```

而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

对于语句 `v.Scale(5)` ，即便 `v` 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 `Scale` 方法有一个指针接收者，为方便起见，Go 会将语句 `v.Scale(5)` 解释为 `(&v).Scale(5)` 。

```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(&p, 8)

	fmt.Println(v, p)
	
	// 和上面的一样
	q := &Vertex{4, 3}
	q.Scale(3)
	ScaleFunc(q, 8)

	fmt.Println(q)
}
```

## 方法与指针重定向（续）

同样的事情也发生在相反的方向。

接受一个值作为参数的函数必须接受一个指定类型的值：

```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！
```

而以值为接收者的方法被调用时，接收者既能为值又能为指针：

```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

这种情况下，方法调用 `p.Abs()` 会被解释为 `(*p).Abs()` 。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
```

## 选择值或指针作为接收者

使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

在本例中， `Scale` 和 `Abs` 接收者的类型为 `*Vertex` ，即便 `Abs` 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。 

```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

模拟面向对象

```go
package puser

struct user {
    int a
    int b
    int c
}

func (this *user) sum() int{
    return this.a + this.b + this.c
}

```
### method方法重写

```go
package main

import "fmt"

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human //匿名字段
    school string
}

type Employee struct {
    Human //匿名字段
    company string
}

//Human定义method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
        e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
    mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
    sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

    mark.SayHi()
    sam.SayHi()
}
```

### 指针类型和unsafe.Pointer

我们一般将 *T 看作指针类型，表示一个指向 T 类型变量的指针。我们都知道 Go 是强类型语言，声明变量之后，变量的类型是不可以改变的，不同类型的指针也不允许相互转化。例如下面这样：

```
 i := 30
 iPtr1 := &i
 var iPtr2 *int64 = (*int64)(iPtr1)
 编译报错：cannot convert iPtr1 (type *int) to type *int64，提示不能进行强制转化。
```

如何实现相互转化

Go 官方提供了 unsafe 包

unsafe.Pointer **通用指针类型**，一种特殊类型的指针，可以包含任意类型的地址，能实现不同的指针类型之间进行转换，类似于 C 语言里的 void* 指针。

```haskell
type ArbitraryType int

type Pointer *ArbitraryType
```

从定义可以看出，Pointer 实际上是 *int。

官方文档里还描述了 Pointer 的四种操作规则：

1. 任何类型的指针都可以转化成 unsafe.Pointer；
2. unsafe.Pointer 可以转化成任何类型的指针；
3. uintptr 可以转换为 unsafe.Pointer；
4. unsafe.Pointer 可以转换为 uintptr；

不同类型的指针允许相互转化实际上是运用了第 1、2 条规则，我们就着例子看下：

```go
func main(){
 i := 30
 iPtr1 := &i

 var iPtr2 *int64 = (*int64)(unsafe.Pointer(iPtr1))

 *iPtr2 = 8

 fmt.Println(i)
}
```

输出：

```undefined
8
```

上面的代码，我们可以把 *int 转为 *int64，并且对新的 *int64 进行操作，从输出会发现 i 的值被改变了。

可以说 unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换。

### uintptr

**参照注释我们知道：**

- uintptr 是一个整数类型（这个非常重要），注意，他不是个指针；
- 但足够保存任何一种指针类型。

**unsafe 包支持了这些方法来完成【类型】=> uintptr 的转换：**

```
func Sizeof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Alignof(x ArbitraryType) uintptr
```

你可以将任意类型变量转入，获取对应语义的 uintptr，用来后续计算内存地址（比如基于一个结构体字段地址，获取下一个字段地址等）。

源码定义：

```vhdl
// uintptr is an integer type that is large enough to hold the bit pattern of
// any pointer.
type uintptr uintptr
```

uintptr 是 Go 内置类型，表示无符号整数，可存储一个完整的地址。常用于指针运算，只需将 unsafe.Pointer 类型转换成 uintptr 类型，做完加减法后，转换成 unsafe.Pointer，通过 * 操作，取值或者修改值都可以。

下面是一个通过指针偏移修改结构体成员的例子，演示下 uintptr 的用法：

```go
type Admin struct {
 Name string
 Age int
}

func main(){
 admin := Admin{
  Name: "seekload",
  Age: 18,
 }
 ptr := &admin
 name := (*string)(unsafe.Pointer(ptr))   // 1

 *name = "四哥"

 fmt.Println(*ptr)

 age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.Age)))  // 2
 *age = 35

 fmt.Println(*ptr)
  
  x := [4]byte{10, 11, 12, 13}
	elPtr := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + 3*unsafe.Sizeof(x[0])))
  fmt.Println(*elPtr)

}
```

输出：

```undefined
{四哥 18}
{四哥 35}
```

特别提下，**unsafe.Offsetof** 的作用是返回成员变量 x 在结构体当中的偏移量，即返回结构体初始内存地址到 x 之间的字节数。

//1 因为**结构体初始地址就是第一个成员的地址**，又 Name 是结构体第一个成员变量，所以此处不用偏移，我们拿到 admin 的地址，然后通过 unsafe.Pointer 转为 *string，再进行赋值操作即可。注意：如果结构体第一个不是Name, 会报错

//2 成员变量 Age 不是第一个字段，想要修改它的值就需要内存偏移。我们先将 admin 的指针转化为 uintptr，再通过 unsafe.Offsetof() 获取到 Age 的偏移量，两者都是 uintptr，进行相加指针运算获取到成员 Age 的地址，最后需要将 uintptr 转化为 unsafe.Pointer，再转化为 *int，才能对 Age 操作。



1. unsafe.Pointer 可以实现不同类型指针之间相互转化；
2. uintptr 搭配着 unsafe.Pointer 使用，实现指针运算；





## 接口

**接口类型** 是由一组方法签名定义的集合。

接口类型的值可以保存任何实现了这些方法的值。

**注意：** 示例代码的 22 行存在一个错误。 由于 `Abs` 方法只为 `*Vertex` （指针类型）定义， 因此 `Vertex` （值类型）并未实现 `Abser` 。

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// cannot use v (type Vertex) as type Abser in assignment:
	// Vertex does not implement Abser (Abs method has pointer receiver)
	// a = v
	
	// 要改成
	a = &v

	fmt.Println(a.Abs())
}


```

## 接口与隐式实现

类型通过实现一个接口的所有方法来实现该接口。 既然无需专门显式声明，也就没有“implements“关键字。

隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

## 接口值

在内部，接口值可以看做包含值和具体类型的元组：

```
(value, type)
```

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 底层值为 nil 的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 `M` 方法）。

**注意：** 保存了 nil 具体值的接口其自身并不为 nil 。

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## nil 接口值

nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 **具体** 方法的类型。

```go
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	//i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 空接口

指定了零个方法的接口值被称为 **空接口：**

```
interface{}
```

空接口可保存任何类型的值。 （因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。 例如，`fmt.Print` 可接受类型为 `interface{}`的任意数量的参数。

```go
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 类型断言



```undefined
x.(T)
```

- 断言x不为nil且x为T类型
- 如果T不是接口类型，则该断言x为T类型
- 如果T类接口类型，则该断言x实现了T接口

### 实例1



```go
func main() {
    var x interface{} = 7
    i := x.(int)
    fmt.Println(reflect.TypeOf(i))
    j := x.(int32)
    fmt.Println(j)
}
```

> 直接赋值的方式，如果断言为true则返回该类型的值，如果断言为false则产生runtime panic；j这里赋值直接panic

输出



```go
int
panic: interface conversion: interface {} is int, not int32

goroutine 1 [running]:
main.main()
        type_assertion.go:12 +0xda
exit status 2
```

不过一般为了避免panic，通过使用ok的方式

```go
func main() {
    var x interface{} = 7
    j, ok := x.(int32)
    if ok {
        fmt.Println(reflect.TypeOf(j))
    } else {
        fmt.Println("x not type of int32")
    }
}
```

### switch type

另外一种就是variable.(type)配合switch进行类型判断



```go
func main() {
    switch v := x.(type) {
    case int:
        fmt.Println("x is type of int", v)
    default:
        fmt.Printf("Unknown type %T!\n", v)
    }
}
```

### 判断struct是否实现某个接口



```go
type shape interface {
    getNumSides() int
    getArea() int
}
type rectangle struct {
    x int
    y int
}

func (r *rectangle) getNumSides() int {
    return 4
}
func (r rectangle) getArea() int {
    return r.x * r.y
}

func main() {
    // compile time Verify that *rectangle implement shape
    var _ shape = &rectangle{}
    // compile time Verify that *rectangle implement shape
    var _ shape = (*rectangle)(nil)

    // compile time Verify that rectangle implement shape
    var _ shape = rectangle{}
  	// 上面会报错，
}
```

输出

```rust
cannot use rectangle literal (type rectangle) as type shape in assignment:
        rectangle does not implement shape (getNumSides method has pointer receiver)
```

### 小结

- `x.(T)`可以在运行时判断x是否为T类型，如果直接使用赋值，当不是T类型时则会产生runtime panic
- 使用`var _ someInterface = someStruct{}`可以在编译时期校验某个struct或者其指针类型是否实现了某个接口

**类型断言** 提供了访问接口值底层具体值的方式。

```
t := i.(T)
```

该语句断言接口值 `i` 保存了具体类型 `T` ，并将其底层类型为 `T` 的值赋予变量 `t` 。

若 `i` 并未保存 `T` 类型的值，该语句就会触发一个panic。

为了 **判断** 一个接口值是否保存了一个特定的类型， 类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

```
t, ok := i.(T)
```

若 `i` 保存了一个 `T` ，那么 `t` 将会是其底层值，而 `ok` 为 `true` 。

否则， `ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，程序并不会产生panic。

请注意这种语法和读取一个映射时的相同之处。

```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

## 类型选择

**类型选择** 是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

```go
switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
```

类型选择中的声明与类型断言 `i.(T)` 的语法相同，只是具体类型 `T` 被替换成了关键字 `type` 。

此选择语句判断接口值 `i` 保存的值类型是 `T` 还是 `S` 。 在 `T` 或 `S` 的情况下，变量 `v` 会分别按 `T` 或 `S` 类型保存 `i` 拥有的值。 在默认（即没有匹配）的情况下，变量 `v` 与 `i` 的接口类型和值相同。

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!
```

## Stringer

[`fmt`](https://go-zh.org/pkg/fmt/) 包中定义的 [`Stringer`](https://go-zh.org/pkg/fmt/#Stringer) 是最普遍的接口之一。

```
type Stringer interface {
    String() string
}
```

`Stringer` 是一个可以用字符串描述自己的类型。`fmt` 包（还有很多包）都通过此接口来打印值。

```go
type Person struct {
	Name string
	Age  int
}

// 相当于定义Person类型的string() 方法，被Println
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

```go
package main

import "fmt"
import "strings"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (p IPAddr) String() string {
	var ret = [] string{} 
	for _, value := range p {
		ret = append(ret, fmt.Sprintf("%d", value))
	}
	return fmt.Sprintf("%s", strings.Join(ret, ","))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

## 错误

Go 程序使用 `error` 值来表示错误状态。

与 `fmt.Stringer` 类似， `error` 类型是一个内建接口：

```
type error interface {
    Error() string
}
```

（与 `fmt.Stringer` 类似， `fmt` 包在打印值时也会满足 `error` 。）

通常函数会返回一个 `error` 值，调用的它的代码应当判断这个错误是否等于 `nil` 来进行错误处理。

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

`error` 为 nil 时表示成功；非 nil 的 `error` 表示失败。

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// 2018-11-08 15:39:37.449579 +0800 CST m=+0.000415854, it didn't work
```

```go
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


func Sqrt(x float64) (float64, error) {
	if(x<0){
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}



func main() {
	v, err := Sqrt((12));
	if  err != nil{
		fmt.Printf("couldn't convert number: %v\n", err)
		return 
	}
	fmt.Println(v)
}
```

## Reader

`io` 包指定了 `io.Reader` 接口， 它表示从数据流的末尾进行读取。

Go 标准库包含了该接口的[许多实现](https://go-zh.org/search?q=Read#Global)， 包括文件、网络连接、压缩和加密等等。

`io.Reader` 接口有一个 `Read` 方法：

```
func (T) Read(b []byte) (n int, err error)
```

`Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。 在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

示例代码创建了一个 [`strings.Reader`](https://go-zh.org/pkg/strings/#Reader) 并以每次 8 字节的速度读取它的输出。

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

## 图像

[`image`](https://go-zh.org/pkg/image/#Image) 包定义了 `Image` 接口：

```go
package image

type Image interface {
	// 颜色模式
    ColorModel() color.Model
    // 图片边界
    Bounds() Rectangle
    // 某个点的颜色
    At(x, y int) color.Color
}
```

**注意：** `Bounds` 方法的返回值 `Rectangle` 实际上是一个 [`image.Rectangle`](https://go-zh.org/pkg/image/#Rectangle)， 它在 `image` 包中声明。

（请参阅[文档](https://go-zh.org/pkg/image/#Image)了解全部信息。）

`color.Color` 和 `color.Model` 类型也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由[`image/color`](https://go-zh.org/pkg/image/color/)包定义。

```go
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
```

## Sync

`WaitGroup` 对象内部有一个计数器，最初从0开始，它有三个方法：`Add(), Done(), Wait()` 用来控制计数器的数量。`Add(n)` 把计数器设置为`n` ，`Done()` 每次把计数器`-1` ，`wait()` 会阻塞代码的运行，直到计数器的值减为`0`。

```go
func main() {
    wg := sync.WaitGroup{}
    wg.Add(100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

这里首先把`wg` 计数设置为`100`， 每个`for`循环运行完毕都把计数器减一，主函数中使用`Wait()` 一直阻塞，直到`wg`为零——也就是所有的`100`个`for`循环都运行完毕。相对于使用管道来说，`WaitGroup` 轻巧了许多。

### 1. 计数器不能为负值

不能使用`Add()` 给`wg` 设置一个负值

使用`Done()` 也不要把计数器设置成负数了。

## Goroutines协程

A *goroutine* is a lightweight thread managed by the Go runtime.

```
go f(x, y, z)
```

会启动一个新的 Go 程并执行

```
f(x, y, z)
```

`f` 、 `x` 、 `y` 和 `z` 的求值发生在当前的 Go 程中，而 `f` 的执行发生在新的 Go 程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。[`sync`](https://go-zh.org/pkg/sync/) 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法

## 信道chan

channel的作用就是在多线程之间传递数据

信道是带有类型的管道，你可以通过它用信道操作符 `<-` 来发送或者接收值。

```go
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
v,ok := <-ch // 如果要知道是否closed得加ok
```

（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

```go
ch := make(chan int)
```

### 单向chan

单向channel变量的声明非常简单，如下：

```go
var ch1 chan int    // ch1是一个正常的channel，不是单向的

var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据

var ch3 <-chan int   // ch3是单向channel，只用于读取int数据

```

单向chan函数参数

```go
func downloadWorker(jobs <-chan *Jobs, results chan<- *Results) {
  for j := range jobs {
		err := Download(j.FileUrl)

		if err != nil {
			logger.Error(err.Error())
		}
		results <- &Results{ID: j.ID}
	}
}
```

可以将 channel 隐式转换为单向队列，只收或只发，不能将单向 channel 转换为普通 channel：

```go
    c := make(chan int, 3)

    var send chan<- int = c // send-only

    var recv <-chan int = c // receive-only
```



### chan基础

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。 一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

```go
package main

import "fmt"
import "math"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
	fmt.Println(math.Pow(2,3),2^3)
	fmt.Println(1<<0)
}
```

## 带缓冲的信道

信道可以是 *带缓冲的* 。将缓冲长度作为第二个参数提供给 `make` 来初始化一个带缓冲的信道：

```go
ch := make(chan int, 100)
```

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

修改示例填满缓冲区，然后看看会发生什么。

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 错误 缓冲区 满了  
    // ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// error  is empty
	fmt.Println(<-ch)

}
```

## range 和 close

发送者可通过 `close` 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

```
v, ok := <-ch
```

之后 `ok` 会被设置为 `false` 。

循环 `for i := range c` 会不断从信道接收值，直到它被关闭。

**注意：** 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

**还要注意：** 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有值需要发送的时候才有必要关闭，例如终止一个 `range` 循环。

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

## select 语句

`select` 语句使一个 Go 程可以等待多个通信操作。

`select` 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

## 默认选择

当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

```go
select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

## chan实例

### channel状态

channel作为go的一种基本数据类型，它有3种基本状态：nil、open、closed：

```go
/* nil channel */
var ch = chan string // A channel is in a nil state when it is declared to its zero value
ch = nil // A channel can be placed in a nil state
 
/* open channel */
ch := make(chan string) // A channel is in a open state when it’s made using the built-in function make.
 
/* closed channel */
close(ch) // A channel is in a closed state when it’s closed using the built-in function close.
```

处于closed状态的channel，执行send操作（ch <- data）将会触发panic异常，而receive操作（<- ch）则不会，这表明了在channel被close之后，goruntine仍然可以从channel取走数据，**如果channel中没有数据可取时，receive操作会立刻返回零值（nil）**。

range循环可以直接在channel上迭代，当channel被关闭并且没有数据时可以直接跳出循环。

另外，对于nil和closed状态的channel执行close操作也会触发panic异常。

#### unbufferd channel和bufferd channel

虽然channel最常用于goroutine之间的通信，但是channel上的send和receive操作并不一定需要携带数据。根据channel是否需要传递数据，可以区分出一些channel的使用场景

**没有数据的channel的使用场景：**

- goroutine A通过channel告诉goroutine B：”请停止正在做的事情“
- goroutine A通过channel告诉goroutine B：”我完成了要做的事情，但是没有任何结果需要反馈“

通知的方式一般是close操作，goroutine A对channel执行了close操作，而goruntine B得到channel已经被关闭这个信息后可以执行相应的处理。使用没有数据的channel的好处：一个goroutine可以同时给多个goroutine发送消息，只是这个消息不携带额外的数据，所以常被用于批量goruntine的退出。

对于这种不携带数据，只是作为信号的channel，一般使用如下：

```Go
ch := make(chan struct{})
ch <- struct{}{}
<- ch
```

**带有数据的channel的使用场景：**

- goroutine A通过channel告诉goroutine B：”请根据我传递给你的数据开始做一件事情“
- goroutine A通过channel告诉goroutine B：”我完成了要做的事情，请接收我传递的数据（结果）“

通知的方式就是goroutine A执行send发送数据，而goroutine B执行receive接收数据。channel携带的数据只能被一个goruntine得到，一个goruntine取走数据后这份数据在channel里就不复存在了。

对于需要携带数据的channel，一般又可以分成带有buffer的channel（bufferd channel）和不带buffer的channel（unbufferd channel）。

**unbufferd channel**

对于unbufferd channel，不存储任何数据，只负责数据的流通，并且数据的接收一定发生在数据发送完成之前。更详细的解释是，goroutine A在往channel发送数据完成之前，一定有goroutine B在等着从这个channel接收数据，否则发送就会导致发送的goruntine被block住，所以发送和接收的goruntine是耦合的。

#### 示例1

```go
// 看下面这个例子，往ch发送数据时就使main gouruntine被永久block住，导致程序死锁。
 	var ch = make(chan string)
 	ch <- "hello" 
//fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan send]:
 	fmt.Println(<-ch)
```

#### 示例2

​	有人可能会考虑将接收操作放到前面，不幸的是仍然导致了死锁，

```go
//  因为channel里没有数据，当前goruntine也会被block住，导致程序死锁。

  var ch = make(chan string)
 	fmt.Println(<-ch) 
 	//fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan receive]:
 	ch <- "hello"
```

#### 示例3

```go
// 在另一个goruntine中执行receive，程序就可以正常工作了。
// 因为在main goruntine发送完数据之前，sub goroutine已经在等待接收数据。
	var ch = make(chan string)
	go func() {
	    fmt.Println("Hello, 世界2")
        fmt.Println(<-ch) //out: hello 
    }()
  ch <- "hello" 
  fmt.Println(time.Millisecond)
```

#### 示例4

    // 我们期望在sub goruntine中打印10个数，实际上却只有main goruntine打印了hello。
    // 因为在sub goruntine打印之前，main goruntine就已经执行完成并退出了
```go
	go func() {
        for i := 0; i < 10; i++ { 
            fmt.Printf("%d ", i)
        }
    }()
    fmt.Println("Hello, 世界3")
```

#### 示例5

这个时候就可以用一个unbufferd channel来让两个goruntine之间有一些通信，让main goruntine收到sub goruntine通知后再退出。在这种场景中，channel并不携带任何数据，只是起到一个信号的作用。

```go
func main() {
	var ch = make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
		}
		ch <- "exit"
	}()
	fmt.Println("hello")
	<-ch
}
```

**bufferd channel**

对带有缓冲区的channel执行send和receive操作，其行为和不带缓冲区的channel不太一样。继续讨论最开始的例子，不过这次的channel是一个size=1的bufferd channel，将数据发送给channel后，数据被拷贝到channel的缓冲区，goruntine继续往后执行，所以程序可以正常工作。

```go
// 这次的channel是一个size=1的bufferd channel，将数据发送给channel后，数据被拷贝到channel的缓冲区，goruntine继续往后执行，所以程序可以正常工作
 	var ch = make(chan string, 1)
 	ch <- "hello"
 	fmt.Println(<-ch) //hello
	
// 	当我们调换发送和接收的顺序时，程序又进入了死锁。因为当channel里没有数据时（干涸），执行receive的goroutine也会被block住，最终导致了死锁
	var ch = make(chan string, 1)
	fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan receive]:
	ch <- "hello"
	
// 	对于buffer size=1的channel，第二个数据发送完成之前，之前发送的第一个数据一定被取走了，否则发送也会被block住
// 对于buffer size>1的channel，发送数据时，之前发送的数据不能保证一定被取走了，并且buffer size越大，数据的交付得到的保证越少
// 根据发送数据、接收数据、数据处理的速度合理的设计buffer size，甚至可以在不浪费空间的情况下做到没有任何延迟
// 如果channel buffer已经塞满了数据，继续执行发送会导致当前goruntine被block住（阻塞），直到channel中的数据被取走一部分才可以继续向channel发送数据
```



## 数据结构与算法



#### 冒泡排序

```go
package main

import (
    "fmt"
    
)

// 冒泡排序函数
func bubbleSort(arr []int) {
    // 获取数组长度
    n := len(arr)
    // 外层循环控制遍历次数
    for i := 0; i < n-1; i++ {
        // 设置一个标志变量
        flag := false
        // 内层循环控制比较和交换
        for j := 0; j < n-i-1; j++ {
            // 如果前一个元素大于后一个元素，就交换它们的位置，并把标志变量设为true
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                flag = true
            }
        }
        // 如果标志变量为false，说明没有发生交换，数组已经有序，直接退出循环
        if !flag {
            break
        }
    }
}

// 测试代码
func main() {
    // 定义一个数组
    arr := []int{5, 3, 7, 2, 9, 4}
    // 打印原始数组
    fmt.Println("Original array:", arr)
    // 调用冒泡排序函数
    bubbleSort(arr)
    // 打印排序后的数组
    fmt.Println("Sorted array:", arr)
}
```





## 等价二叉树

不同二叉树的叶节点上可以保存相同的值序列

在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。 我们将使用 Go 的并发和信道来编写一个简单的解法。

```


本例使用了 tree 包，它定义了类型：

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

**1.** 实现 `Walk` 函数。

**2.** 测试 `Walk` 函数。

函数 `tree.New(k)` 用于构造一个随机结构的二叉树，它保存了值 `k` 、 `2k` 、 `3k`... `10k` 。

创建一个新的信道 `ch` 并且对其进行步进：

```
go Walk(tree.New(1), ch)
```

然后从信道中读取并打印 10 个值。应当是数字 `1, 2, 3, ..., 10` 。

**3.** 用 `Walk` 实现 `Same` 函数来检测 `t1` 和 `t2` 是否存储了相同的值。

**4.** 测试 `Same` 函数。

`Same(tree.New(1), tree.New(1))` 应当返回 `true` ，而 `Same(tree.New(1), tree.New(2))` 应当返回 `false` 。

`Tree` 的文档可在[这里](https://godoc.org/golang.org/x/tour/tree#Tree)找到

```go
package main

import (
"golang.org/x/tour/tree"
"fmt"
)
//  发送value，结束后关闭channel
func Walk(t *tree.Tree, ch chan int){
    sendValue(t,ch)
    close(ch)
}
//  递归向channel传值
func sendValue(t *tree.Tree, ch chan int){
    if t != nil {
        sendValue(t.Left, ch)
        ch <- t.Value
        sendValue(t.Right, ch)
   }
}

// 使用写好的Walk函数来确定两个tree对象  是否一样 原理还是判断value值
func Same(t1, t2 *tree.Tree) bool {
       ch1 := make(chan int)
       ch2 := make(chan int)
       go Walk(t1,ch1)
       go Walk(t2,ch2)
       for i := range ch1 {   // ch1 关闭后   for循环自动跳出
               if i != <- ch2 {
                      return false
               }
       }
      return true
}

func main() {

    // 打印 tree.New(1)的值
    var ch = make(chan int)
    go Walk(tree.New(1),ch)
    for v := range ch {
          fmt.Println(v)
    }
    
    var ch2 = make(chan int)
    go Walk(tree.New(2),ch2)
    for v := range ch2 {
          fmt.Println(v)
    }

    //  比较两个tree的value值是否相等
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
```

## sync.Mutex

我们已经看到信道非常适合在各个 Go 程间进行通信。

但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？

这里涉及的概念叫做 *互斥 （mutual exclusion）* ，我们通常使用 *互斥锁 （Mutex）*这一数据结构来提供这种机制。

Go 标准库中提供了 [`sync.Mutex`](https://go-zh.org/pkg/sync/#Mutex) 互斥锁类型及其两个方法：

- `Lock`
- `Unlock`

我们可以通过在代码前调用 `Lock` 方法，在代码后调用 `Unlock` 方法来保证一段代码的互斥执行。 参见 `Inc` 方法。

我们也可以用 `defer` 语句来保证互斥锁一定会被解锁。参见 `Value` 方法。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

### 练习 Web 爬虫

```go
package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type result struct {
	url, body string
	urls []string
	err error
	depth int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	results := make(chan *result)
	fetched := make(map[string]bool)
	fetch   := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		results <- &result{url, body, urls, err, depth}
	}

	go fetch(url, depth)
	fetched[url] = true

	// 1 url is currently being fetched in background, loop while fetching
	for fetching := 1; fetching > 0; fetching-- {
		res := <- results

		// skip failed fetches
		if res.err != nil {
			fmt.Println(res.err)
			continue
		}

		fmt.Printf("found: %s %q\n", res.url, res.body)

		// follow links if depth has not been exhausted
		if res.depth > 0 {
			for _, u := range res.urls {
				// don't attempt to re-fetch known url, decrement depth
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth - 1)
					fetched[u] = true
				}
			}
		}
	}

	close(results)
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls     []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
```

### 连接数据库



## 接下来去哪？

[Go 文档](https://go-zh.org/doc/) 是一个极好的 开始。 它包含了参考、指南、视频等等更多资料。

了解如何组织 Go 代码并在其上工作，参阅 [此视频](http://www.youtube.com/watch?v=XCsL89YtqCs)，或者阅读 [如何编写 Go 代码](https://go-zh.org/doc/code.html)。

如果你需要标准库方面的帮助，请参考 [包手册](https://go-zh.org/pkg/)。如果是语言本身的帮助，阅读[语言规范](https://go-zh.org/ref/spec)是件令人愉快的事情。

进一步探索 Go 的并发模型，参阅 [Go 并发模型](http://www.youtube.com/watch?v=f6kdp27TYZs) ([幻灯片](http://talks.go-zh.org/2012/concurrency.slide)) 以及 [深入 Go 并发模型](https://www.youtube.com/watch?v=QDDwwePbDtw) ([幻灯片](http://talks.go-zh.org/2013/advconc.slide)) 并阅读 [通过通信共享内存](https://go-zh.org/doc/codewalk/sharemem/) 的代码之旅。

想要开始编写 Web 应用，请参阅 [一个简单的编程环境](http://vimeo.com/53221558) ([幻灯片](http://talks.go-zh.org/2012/simple.slide)) 并阅读 [编写 Web 应用](https://go-zh.org/doc/articles/wiki/)的指南.

[函数：Go 中的一等公民](https://go-zh.org/doc/codewalk/functions/)展示了有趣的函数类型。

[Go 博客](http://blog.go-zh.org/) 有着众多关于 Go 的文章和信息。

[mikespook 的博客](http://www.mikespook.com/tag/golang/)中有大量中文的关于 Go 的文章和翻译。

开源电子书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang) 和 [Go 入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN) 能够帮助你更加深入的了解和学习 Go 语言。

访问 [go-zh.org](https://go-zh.org/) 了解更多内容。

### Go包依赖管理

Go dependency management summary:

- `vgo` if your go version is: `x >= go 1.11`
- `dep` or `vendor` if your go version is: `go 1.6 >= x < go 1.11`
- Manually if your go version is: `x < go 1.6`

------

Go 1.11 has a feature `vgo` which will [replace](https://github.com/golang/go/wiki/vgo#current-state) `dep`.

To use `vgo`, see [Modules](https://github.com/golang/go/wiki/Modules) documentation. TLDR below:

```golang
export GO111MODULE=on
go mod init
go mod vendor # if you have vendor/ folder, will automatically integrate
go build
```

This method creates a file called `go.mod` in your projects directory. You can then build your project with `go build`. If `GO111MODULE=auto` is set, then your project cannot be in `$GOPATH`.

### Go模块的使用

Go模块是Go语言的依赖包管理工具。

1、Go1.11及以后版本才能使用。 
2、Go1.11需要设置环境变量 GO111MODULE 为 on（新特性开关，按照Go语言惯例，mod首次在go1.11版本中使用，go1.12及以后版本这个设置应该不会用了）。

```go
GO111MODULE=on  go run .
```

mod是模块英文modules的简写。

#### GO111MODULE 环境变量

`GO111MODULE`可以设置为三个字符串值之一：off，on或auto（默认值）。

- off，则go命令从不使用新模块支持。它查找vendor 目录和GOPATH以查找依赖关系;也就是继续使用“GOPATH模式”。
- on，则go命令需要使用模块，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod下载依赖。
- auto或未设置，则go命令根据当前目录启用或禁用模块支持。仅当当前目录位于GOPATH/src之外并且其本身包含go.mod文件或位于包含go.mod文件的目录下时，才启用模块支持。

#### 常用的命令行：

- `go help mod`查看帮助。
- `go mod init <项目模块名称>`初始化模块，会在项目根目录下生成 `go.mod` 文件。参数`<项目模块名称>`是非必写的，但如果你的项目还没有代码编写，这个参数能快速初始化模块。如果之前使用其它依赖管理工具(比如dep，glide等)，mod会自动接管原来依赖关系。
- `go mod tidy`根据go.mod文件来处理依赖关系。
- `go mod vendor`将依赖包复制到项目下的 vendor 目录。建议一些使用了被墙包的话可以这么处理，方便用户快速使用命令`go build -mod=vendor`编译。
- `go list -m all`显示依赖关系。`go list -m -json all`显示详细依赖关系。
- `go mod download <path@version>`下载依赖。参数`<path@version>`是非必写的，path是包的路径，version是包的版本。
- 其它命令可以通过`go help mod`来查看。

go.mod文件是文本文件，是可以自己手动编辑的。 
Go模块版本控制的下载文件及信息会存储到GOPATH的pkg/mod文件夹里。 
使用了Go模块，源码不一定要在GOPATH中进行。

go.mod文件必须要提交到git仓库
问：启用Go模块以后，使用go get xxx时会报错提示"go: cannot find main module; see 'go help modules'"，这个是怎么回事？
答：因为没有找到go.mod文件，所以会报错。你只要在项目根目录下生成一个go.mod文件就可以了。

问：如何在Go模块里使用本地依赖包？
答：首先在项目的go.mod文件的require处添加依赖包，然后在replace处添加替换本地依赖包(路径要处理妥当)。比如：

require (
​    mytest v0.0.0
)
replace (
​    mytest v0.0.0 => ../mytest
)

#### Go模块参考资料

- 语义化版本(中文) <https://semver.org/lang/zh-CN/>
- Go模块官方文档(英文) [https://github.com/golang/go/...](https://github.com/golang/go/wiki/Modules)
- Go模块命令说明(英文) [https://golang.google.cn/cmd/...](https://golang.google.cn/cmd/go/#hdr-Module_maintenance)

### Go Test

#### 运行测试

```
$ go test .          # Run all tests in the current directory
$ go test ./...      # Run all tests in the current directory and sub-directories
$ go test ./foo/bar  # Run all tests in the ./foo/bar directory
```

#### 分析测试覆盖率

```
 go test -cover ./...
```

#### 测试所有依赖

在为发布或者部署可执行文件，或者是公开分发代码之前，你可能会想要运行 `go test all` 命令：

```
$ go test all
```

这个命令将会对模块中的所有包和所有依赖项运行测试（包括测试*测试依赖*以及必要的*标准库包*）。并且它能够帮助验证所使用的依赖项的确切版本是否彼此兼容

#### 压力测试

可以使用 `go test -count` 命令来连续多次运行一个测试。这在想要检查偶发或者间歇性失败的时候很有用

```
$ go test -run=^TestFooBar$ -count=500 .
```



### 执行静态分析

`go vet` 工具会对你的代码进行静态分析，然后对_可能_出现（编译器可能无法获取）的代码错误进行示警。像是无法访问的代码、不必要的赋值和错误格式的构建标记等问题。可以像这样使用它：

```
$ go vet foo.go     # Vet the foo.go file
$ go vet .          # Vet all files in the current directory
$ go vet ./...      # Vet all files in the current directory and sub-directories
$ go vet ./foo/bar  # Vet all files in the ./foo/bar directory
```



### GoDoc的使用

```
$ go doc strings            # View simplified documentation for the strings package
$ go doc -all strings       # View full documentation for the strings package
$ go doc strings.Replace    # View documentation for the strings.Replace function
$ go doc sql.DB             # View documentation for the database/sql.DB type
$ go doc sql.DB.Query       # View documentation for the database/sql.DB.Query method
```

还可以使用 `-src` 标志来展示相关的 Go 源代码。例如：

```
$ go doc -src strings.Replace   # View the source code for the strings.Replace function
```

#### 一. 约定

1. 注释符`//`后面要加空格, 例如: `// xxx`

2. 在`package, const, type, func`等`关键字`上面并且紧邻关键字的注释才会被展示

   ```go
   // 此行注释被省略
   
   // 此行注释被展示 
   // 
   // 此行注释被展示2 
   package banana
   ```

3. `type, const, func`以**名称**为注释的开头, `package`以`Package name`为注释的开头

   ```go
   // Package banana ...
   package banana
   
   // Xyz ...
   const Xyz = 1
   
   // Abc ...
   type Abc struct {}
   
   // Bcd ...
   func Bcd() {}
   ```

4. 有效的关键字注释不应该超过`3`行

   ```
   // Package banana ...
   // ...
   // ...
   // 最好不要超过三行
   package banana
   ```

5. `Package`的注释如果超过`3`行, 应该放在当前包目录下一个单独的文件中, 如:doc.go

6. **如果当前包目录下包含多个Package注释的go文件(包括doc.go), 那么按照文件名的字母数序优先显示**

   ```
   //----- doc.go -----
   
   // Package banana ...第一个显示
   package banana
   ```

   ```
   //----- e.go -----
   
   // Package banana ...第二个显示
   package banana
   ```

   ```
   //----- f.go -----
   
   // Package banana ...第三个显示
   package banana
   ```

7. `Package`的注释会出现在godoc的[包列表](https://golang.org/pkg/)中, 但只能展示大约523字节的长度

8. 在无效注释中以`BUG(xxx)`开头的注释, 将被识别为已知bug, 显示在`bugs`区域, [示例](http://golang.org/pkg/bytes/#bugs)

   ```
   // BUG(who): 我是bug说明
   
   // Package banana ...
   package banana
   ```

9. 如果`bug注释`和`关键字注释`中间无换行, 那么`混合的注释`将被显示在`bugs`和`godoc列表`两个区域内

   ```
   // BUG(xxx): 我是bug注释
   // Package banana ...也是pkg注释
   package banana
   ```

10. URL将被转化为HTML链接

#### 二. Example

文件必须放在当前包下

文件名以`example`开头, `_`连接, `test`结尾, 如:`example_xxx_test.go` 

包名是`当前包名` + `_test`, 如: `strings_test` 

函数名称的格式`func Example[FuncName][_tag]()` 

函数注释会展示在页面上

函数结尾加上`// Output:`注释, 说明函数返回的值

### Go make 和 new 的区别

#### new(T) 返回的是 T 的指针

new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值。

```go
p1 := new(int)
fmt.Printf("p1 --> %#v \n ", p1) //(*int)(0xc42000e250) 
fmt.Printf("p1 point to --> %#v \n ", *p1) //0
var p2 *int
i := 0
p2 = &i
fmt.Printf("p2 --> %#v \n ", p2) //(*int)(0xc42000e278) 
fmt.Printf("p2 point to --> %#v \n ", *p2) //0
```

上面的代码是等价的，new(int) 将分配的空间初始化为 int 的零值，也就是 0，并返回 int 的指针，这和直接声明指针并初始化的效果是相同的。

#### make 只能用于 slice,map,channel

make 只能用于 slice，map，channel 三种类型，make(T, args) 返回的是初始化之后的 T 类型的值，这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。

```go
var s1 []int
if s1 == nil {
    fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
}
s2 := make([]int, 3)
if s2 == nil {
    fmt.Printf("s2 is nil --> %#v \n ", s2)
} else {
    fmt.Printf("s2 is not nill --> %#v \n ", s2)// []int{0, 0, 0}
}
```

slice 的零值是 nil，使用 make 之后 slice 是一个初始化的 slice，即 slice 的长度、容量、底层指向的 array 都被 make 完成初始化，此时 slice 内容被类型 int 的零值填充，形式是 [0 0 0]，map 和 channel 也是类似的。

```go
var m1 map[int]string
if m1 == nil {
    fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
}

m2 := make(map[int]string)
if m2 == nil {
    fmt.Printf("m2 is nil --> %#v \n ", m2)
} else {
    fmt.Printf("m2 is not nill --> %#v \n ", m2) map[int]string{} 
}


var c1 chan string
if c1 == nil {
    fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
}

c2 := make(chan string)
if c2 == nil {
    fmt.Printf("c2 is nil --> %#v \n ", c2)
} else {
    fmt.Printf("c2 is not nill --> %#v \n ", c2)//(chan string)(0xc420016120)
}
```

#### make(T, args) 返回的是 T 的 引用

如果不特殊声明，go 的函数默认都是按值传参，即通过函数传递的参数是值的副本，在函数内部对值修改不影响值的本身，但是 make(T, args) 返回的值通过函数传递参数之后可以直接修改，即 map，slice，channel 通过函数穿参之后在函数内部修改将影响函数外部的值。

### struct 和 map

#### query-result-to-map-in-golang

```go
rows, _ := db.Query("SELECT ...") // Note: Ignoring errors for brevity
for rows.Next() {
    m := make(map[string]interface{})
    
    // This WON'T WORK
    if err := rows.Scan(&m); err != nil {
        // ERROR: sql: expected X destination arguments in Scan, not 1
    }
}
// 上面的出错
```

Basically the SQL package is thinking that you expect a single column to be returned and for it to be a **map[string]interface{}** compatible type, which isn’t what we we’re trying to do.

```go
rows, _ := db.Query("SELECT ...") // Note: Ignoring errors for brevity
cols, _ := rows.Columns()

for rows.Next() {
    // Create a slice of interface{}'s to represent each column,
    // and a second slice to contain pointers to each item in the columns slice.
    columns := make([]interface{}, len(cols))
    columnPointers := make([]interface{}, len(cols))
    for i, _ := range columns {
        columnPointers[i] = &columns[i]
    }
    
    // Scan the result into the column pointers...
    if err := rows.Scan(columnPointers...); err != nil {
        return err
    }

    // Create our map, and retrieve the value for each column from the pointers slice,
    // storing it in the map with the name of the column as the key.
    m := make(map[string]interface{})
    for i, colName := range cols {
        val := columnPointers[i].(*interface{})
        m[colName] = *val
    }
    
    // Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...] 
    fmt.Print(m)
}
```

### range

range 是 Golang 语言定义的一种语法糖迭代器，1.5版本 Golang 引入自举编译器后 range 相关源码如下，根据类型不同进行不同的处理，支持对切片和数组、map、通道、字符串类型的的迭代。编译器会对每一种 range 支持的类型做专门的 “语法糖还原”。

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素

 range 遍历开始前会创建副本



数组

```go
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5} 
	 var r [5]int
	 for i, v := range a { 
		if i == 0 {
			a[1] = 12
			a[2] = 13 
		}
		r[i] = v 
	 }
	 fmt.Println("r = ", r) 
	 fmt.Println("a = ", a)
   }
   
   r =  [1 2 3 4 5]
   a =  [1 12 13 4 5]
```

切片

```
package main

import "fmt"

func main() {
    var a = []int{1, 2, 3, 4, 5} 
    var r = make([]int,5)
    for i, v := range a { 
        if i == 0 {
            a[1] = 12
            a[2] = 13 
        }
        r[i] = v 
     }
     fmt.Println("r = ", r) 
     fmt.Println("a = ", a)
   
   }
   //注意变化
   r =  [1 12 13 4 5]
   a =  [1 12 13 4 5]
```

循环过程中依然创建了原切片的副本，但是因为切片自身的结构，创建的副本依然和原切片共享底层数组，只要没发生扩容，他们的值发生变化时就是同步变化的

到这里我们一起来看下遍历数组和切片时源码是什么样的吧？源码比较长，我们大概挑选出来关键的简单汇总就是如下

```
ha := a   //创建副本
hv1 := 0
hn := len(ha)   //循环前长度已经确定
v1 := hv1       //索引变量和取值变量都只在开始时声明，后面都是复用
v2 := nil       
for ; hv1 < hn; hv1++ {
    tmp := ha[hv1]  
    v1, v2 = hv1, tmp
    ...
}
```

这里给的是分析使用 for i, elem := range a {} 遍历数组和切片，同时关心索引和数据的情况，只关心索引或者只关心数据值的代码稍微不同，也就是关不关心 v1 和 v2 ，不关心直接nil掉。

Golang 1.5版本之前的 gcc 源码中语法糖扩展的 range 源码我们也贴出来方便大家理解。

```
// The loop we generate:
//   for_temp := range    //创建副本，数组的话重新复制新数组，切片的话复制新切片后，副本切片与原切片共享底层数组
//   len_temp := len(for_temp)  //循环前长度已经确定
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
```

仔细看这两段代码，原来玄机都藏在这里了~~

#### 循环次数在循环开始前已经确定

循环开始前先计算了数组和切片的长度，for 循环用这个长度来限制循环次数的，也就是循环次数在循环开始前就已经确定了呐，so 循环中再怎么追加或者删除元素都不会影响循环次数，也就不会死循环了~~

```
func main() {
   v := []int{1, 2, 3} 
   counter := 0
   for i := range v {
      counter++
      v = append(v, i) 
   }
   fmt.Println(counter)   //counter代表循环次数，3次哦，没有死循环，也不是6次，虽然v其实已经是长度为6的切片
   fmt.Println(v)   //[1,2,3,0,1,2]
}
```

#### 循环的时候会创建每个元素的副本

```
package main

import "fmt"

type T struct {
     n int
 }
 func main() {
     ts := [2]T{}
     for i, t := range ts {
         switch i {
         case 0:
             t.n = 3
             ts[1].n = 9 
         case 1:
             fmt.Print(t.n, " ") 
         }
     }
     fmt.Print(ts)
 }
```

for-range 循环数组时使用的是数组 ts 的副本，所以 t.n = 3 的赋值操作不会影响原数组。但 `ts[1].n = 9`这种方式操作的确是原数组的元素值，所以是会发生变化的。这也是我们推崇的方法。

### 循环的时候短声明只会在开始时执行一次，后面都是重用

循环 index 和 value 在每次循环体中都会被重用，而不是新声明。for-range 循环里的短声明`index,value :=`相当于第一次是 := ，后面都是 =，所以变量地址是不变的，就相当于全局变量了。

每次遍历会把被循环元素当前 key 和值赋给这两个全局变量，但是注意变量还是那个变量，地址不变，所以如果用的是地址的或者当前上下文环境值的话最后打印出来都是同一个值。



```
package main

import "fmt"

func main() {
     slice := []int{0,1,2,3}
     m := make(map[int]*int)
     for key,val := range slice {
       m[key] = &val
       fmt.Println(key,&key)
       fmt.Println(val,&val)
     }
     for k,v := range m {
      fmt.Println(k,"->",*v)
     }
 }
 
0 0xc000018058
0 0xc000018060
1 0xc000018058
1 0xc000018060
2 0xc000018058
2 0xc000018060
3 0xc000018058
3 0xc000018060
0 -> 3
1 -> 3
2 -> 3
3 -> 3

```

key0、key1、key2、key3 其实都是短声明中的key变量，所以地址是一致的，val0、val1、val2、val3 其实都是短声明中的val变量，地址也一致

最终遍历 map 进行输出时因为 map 赋值时用的是 val 的地址`m[key] = &val`,循环结束时 val 的值是3，所以最终输出时4个元素的值都是3。 

这里需要注意 map 的遍历输出结果 key 的顺序可能会不一致，比如2，0，1，3这样，那是因为 map 的遍历输出是无序的，后面会再说，但是对应的 value 的值都是3。

那如果想要新生成的map也输出正确的值怎么做呐？



```
package main

import "fmt"

func main() {
     slice := []int{0,1,2,3}
     m := make(map[int]*int)
     for key,val := range slice {
       value := val    //增加临时变量，每次都是新声明的，地址也就不一样，也就能传过去正确的值
       m[key] = &value
       fmt.Println(key,&key)
       fmt.Println(val,&val)
     }
     for k,v := range m {
      fmt.Println(k,"->",*v)
     }
 }
 0 0xc0000b2008
0 0xc0000b2010
1 0xc0000b2008
1 0xc0000b2010
2 0xc0000b2008
2 0xc0000b2010
3 0xc0000b2008
3 0xc0000b2010
1 -> 1
2 -> 2
3 -> 3
0 -> 0

```

再来看下 for-range 循环中开启了协程会怎么样？

```
package main

import "fmt"

import "time"
func main() {
      var m = []int{1, 2, 3}
     for i, v := range m {
         go func() {
             fmt.Println(i, v) 
         }()
     }
     time.Sleep(time.Second * 1) 
 }
 
2 3
2 3
2 3
```

各个 goroutine 中输出的 i、v 值都是 for-range 循环结束后的 i、v 最终值，而不是各个 goroutine 启动时的 i, v值。因为 goroutine 执行是在后面的某一个时间，使用的是执行时上下文环境的变量值，i，v又相当于一个全局变量，协程执行时 for-range 循环已结束，i 和 v 都是最后一次循环的值2和3，所以最后输出都是2和3。

两种方法，一种是**临时变量存储循环iv值进行使用**，另外一种是**通过函数参数进行传递** `go func(i,v){}(i,v)`

```
for i, v := range m {
     index := i // 这里的 := 会新声明变量，而不是重用 
     value := v
     go func() {
        fmt.Println(index, value) 
     }()
}
```

```
for i, v := range m { 
    go func(i,v int) {
      fmt.Println(i, v) 
    }(i,v)
}
```

至于 for-range 中通过 append 函数为切片追加元素继而在循环外打印切片时元素值是否发生变化，取决于切片 append 的原理，容量是否足够，是否发生扩容生成新的底层数组，底层数组值是否发生改变等，不是本文的重点，这里就不详细说了~~





```
package main
 
import (
	"fmt"
)
 
func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)
 
	for _,v :=range slice{
		if v==1 {
			v=100
		}
	}
	for k,v :=range slice{
		fmt.Println("k:",k,"v:",v)
	}
}
预想的结果应该是：

k: 0 v: 0
k: 1 v: 100
k: 2 v: 2
k: 3 v: 3

但是实际上

k: 0 v: 0
k: 1 v: 1
k: 2 v: 2
k: 3 v: 3
```

值并没有改变，出现上述问题的原因是因为for range遍历的内容是对原内容的一个拷贝，所以不能用来修改原切片中内容。

使用 k根据索引直接修改值。

```
for k,v :=range slice{
		if v==1 {
			slice[k]=100
		}
	}

```





map

```
 package main
     
    import (
    	"fmt"
    )
     
    func main() {
     
    	s :=[]int{1,2,3,4}
    	m :=make(map[int]*int)
     
    	for k,v:=range s{
    		m[k]=&v
    	}
    	for key, value := range m {
    		fmt.Printf("map[%v]=%v\n", key, *value)
    	}
     
    	fmt.Println(m)
    }
    
    预期打印的值应该为：
map[0]=1
map[1]=2
map[2]=3
map[3]=4

实际结果：
map[2]=4
map[3]=4
map[0]=4
map[1]=4
map[0:0xc00012a008 1:0xc00012a008 2:0xc00012a008 3:0xc00012a008]
```

这里需要注意的是，`for range map`返回的`K-V`键值对顺序是不固定的，是随机的，

从上面结果我们可以猜想到，range指向的都是同一个指针。

其实还是因为for range创建的是每个元素的拷贝，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，值的地址总是相同的。

```
声明一个中间变量，保存value,并且复制给map即可
for k,v:=range s{
		n:=v
		m[k]= &n
	}
```



### 指针易错

```go
// 这种声明方式 p 是一个 nil 值
var p *Point

// 改为
var p *Point = new(Point)

// 或者
var p *Point = &Point{}
```

```go
type Point struct {
    X, Y float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
    var p *Point
    fmt.Println(p.Abs())
}
```



https://github.com/bobvanluijt/golang-map-vs-struct-benchmark



### go build -tags

#### 使用方法

1. 构建约束以一行`+build`开始的注释。在`+build`之后列出了一些条件，在这些条件成立时，该文件应包含在编译的包中；
2. 约束可以出现在任何源文件中，不限于go文件；
3. `+build`必须出现在`package`语句之前，`+build`注释之后应要有一个空行。

 

```go
// 
// +build debug

package main

import "fmt"

func main() {
 fmt.Println("Hello World!")
}
```

#### 语法规则

1）只允许是字母数字或_

2）多个条件之间，空格表示OR；逗号表示AND；叹号(!)表示NOT

3）一个文件可以有多个+build，它们之间的关系是AND。如：



```go
// +build linux darwin
// +build 386
等价于
// +build (linux OR darwin) AND 386
```

 

### 教程

https://www.imooc.com/video/16869

https://book.eddycjy.com/golang/gin/reload-http.html

### 笔记

Go 程序热编译工具，提升开发效率

https://github.com/silenceper/gowatch/blob/master/README_ZH_CN.md



goland 配置SDK时会出现 "The selected directory is not a valid home for Go SDK" 的错误提示

修改 Go path 路径下 C:\Program Files\Go\src\runtime\internal\sys\zversion.go文件, 添加一行版本信息，改成实际的版本号，然后重启goland

```
const TheVersion = `go1.17.6`
```

