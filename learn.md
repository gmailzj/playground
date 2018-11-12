

### 在线书籍教程

[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md)

[Go语言教程](http://www.runoob.com/go/go-data-types.html)

https://gobyexample.com/regular-expressions

### Go 命令基础

#### go env

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

#### go build

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

### GO 语法基础

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

程序中可能会使用到这些标点符号：.、,、;、: 和 …。

#### 标识符

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

#### 关键字

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



### Go语言数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

Go 语言按类别有以下几种数据类型：

| 序号   | 类型和描述                                    |
| ---- | ---------------------------------------- |
| 1    | **布尔型**布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
| 2    | **数字类型**整型 int 和浮点型 float，Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型:**字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本。 |
| 4    | **派生类型:**包括：(a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型 |

#### 数字类型

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

### 约定

fmt.Printf("%格式")

General:

```go
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
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



### Go语言基础语法

#### 变量的初始化

变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```

## 短变量声明

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

## 类型转换

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

与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换

## 类型推导

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

## 常量

常量的声明与变量类似，只不过是使用 `const` 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 `:=` 语法声明



## for

Go 只有一种循环结构： `for` 循环。

基本的 `for` 循环由三部分组成，它们用分号隔开：

- 初始化语句：在第一次迭代前执行
- 条件表达式：在每次迭代前求值
- 后置语句：在每次迭代的结尾执行

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

## for 是 Go 中的 “while”

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

## if

Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 `( )` ，而大括号 `{ }` 则是必须的。

```go
if x < 0 {
	return sqrt(-x) + "i"
}
```

## if 的简短语句

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

## if 和 else

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

## switch

你大概已经知道 `switch` 语句的样子了。

除非以 `fallthrough` 语句结束，否则分支会自动终止。

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

## switch 的求值顺序

switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

（例如，

```go
switch i {
case 0:
case f():
}
```

在 `i==0` 时 `f` 不会被调用。）

## 没有条件的 switch

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

## defer

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

## defer 栈

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

## 结构体

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

## 结构体字段

结构体字段使用点号来访问。

## 结构体指针

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

## 数组

类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

表达式

```
var a [10]int
```

会将变量 `a` 声明为拥有有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是个限制，不过没关系， Go 提供了更加便利的方式来使用数组。

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

## 切片

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

## 切片就像数组的引用

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

## 切片文法

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
## 切片的默认行为

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

## 切片的长度与容量

切片拥有 **长度** 和 **容量** 。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

你可以通过重新切片来扩展一个切片。

## nil 切片

切片的零值是 `nil` 。

nil 切片的长度和容量为 0 且没有底层数组。

## 用 make 创建切片

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

## 切片的切片

切片可包含任何类型，甚至包括其它的切片。

## 向切片追加元素

为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 `append` 函数。 内建函数的[文档](https://go-zh.org/pkg/builtin/#append)对此函数有详细的介绍。

```
func append(s []T, vs ...T) []T
```

`append` 的第一个参数 `s` 是一个元素类型为 `T` 的切片， 其余类型为 `T` 的值将会追加到该切片的末尾。

`append` 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 `s` 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。 返回的切片会指向这个新分配的数组。

（要了解关于切片的更多内容，请阅读文章[Go 切片：用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals)。）

## Range

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

## 映射(Map)

映射的创建的4种方式

make ( map [KeyType] ValueType, initialCapacity )

make ( map [KeyType] ValueType )

map [KeyType ] ValueType {}

map [KeyType ] ValueType { key1 : value1, key2: value2, ... , keyN : valueN}

映射将键映射到值。

映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。

`make` 函数会返回给定类型的映射，并将其初始化备用。

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

Go 没有类。不过你可以为结构体类型定义方法。

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

## 指针接收者

你可以为指针接收者声明方法。

这意味着对于某类型 `T` ，接收者的类型可以用 `*T` 的文法。 （此外， `T` 不能是像 `*int` 这样的指针。）

例如，这里为 `*Vertex` 定义了 `Scale` 方法。

指针接收者的方法可以修改接收者指向的值（就像 `Scale` 在这做的）。 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

试着移除第 16 行 `Scale` 函数声明中的 `*` ，观察此程序的行为如何变化。

若使用值接收者，那么 `Scale` 方法会对原始 `Vertex` 值的副本进行操作。 （对于函数的其它参数也是如此。） `Scale` 方法必须用指针接受者来更改 `main` 函数中声明的 `Vertex` 的值。

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

## 接口

**接口类型** 是由一组方法签名定义的集合。

接口类型的值可以保存任何实现了这些方法的值。

**注意：** 示例代码的 22 行存在一个错误。 由于 `Abs` 方法只为 `*Vertex` （指针类型）定义， 因此 `Vertex` （值类型）并未实现 `Abser` 。

```
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
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

```
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

**类型断言** 提供了访问接口值底层具体值的方式。

```
t := i.(T)
```

该语句断言接口值 `i` 保存了具体类型 `T` ，并将其底层类型为 `T` 的值赋予变量 `t` 。

若 `i` 并未保存 `T` 类型的值，该语句就会触发一个恐慌。

为了 **判断** 一个接口值是否保存了一个特定的类型， 类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

```
t, ok := i.(T)
```

若 `i` 保存了一个 `T` ，那么 `t` 将会是其底层值，而 `ok` 为 `true` 。

否则， `ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，程序并不会产生恐慌。

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

```
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

```
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
v,ok := <-ch // 如果要知道是否closed得加ok
```

（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

```
ch := make(chan int)
```

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

```
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

### Go模块的使用

Go模块是Go语言的依赖包管理工具。

1、Go1.11及以后版本才能使用。 
2、Go1.11需要设置环境变量 GO111MODULE 为 on（新特性开关，按照Go语言惯例，mod首次在go1.11版本中使用，go1.12及以后版本这个设置应该不会用了）。

```go
GO111MODULE=on  go run .
```

mod是模块英文modules的简写。

常用的命令行：

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
    mytest v0.0.0
)
replace (
    mytest v0.0.0 => ../mytest
)

##### 参考资料

- 语义化版本(中文) <https://semver.org/lang/zh-CN/>
- Go模块官方文档(英文) [https://github.com/golang/go/...](https://github.com/golang/go/wiki/Modules)
- Go模块命令说明(英文) [https://golang.google.cn/cmd/...](https://golang.google.cn/cmd/go/#hdr-Module_maintenance)

