### 在线书籍教程

[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md)

[Go语言教程](http://www.runoob.com/go/go-data-types.html)

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