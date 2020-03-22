一、基本数据类型之间的转换

1、string到int  

int,err:=strconv.Atoi(string)  

2、string到int64  

int64, err := strconv.ParseInt(string, 10, 64)  

3、int到string  

string:=strconv.Itoa(int)  

4、int64到string  

string:=strconv.FormatInt(int64,10)  

5、字符串到float32/float64

float32, err = ParseFloat(string, 32)  

float64,err = ParseFloat(string,64)

6、int64转int

int:=int(int64)  

7、int转int64

int64:=int64(int)



二 interface转换

1、interface{}类型转换成具体类型

原理：断言实现。如：

断言成功返回true,失败返回false

var a interface{} = "a

value, ok := a.(string)
if !ok {
    fmt.Println("It's not ok for type string")
    return
}
fmt.Println("The value is ", value)

在Go里面分为命名类型(named，所有使用type定义的类型都是命名类型，如int int64 string bool)和非命名类型(unamed map slice array)，一个非命名类型可以赋值给一个命名类型，只要他们的结构相同



普通变量类型**int,float,string** 都可以使用 `type (a)`这种形式来进行强制类型转换,比如

```
var a int32  = 10
var b int64 = int64(a)
var c float32 = 12.3
var d float64 =float64(c)
```

golang中 指针也是有类型的,

```
package main

func main() {
    var a int = 10
    var p *int =&a
    var c *int64 
    c= (*int64)(p)
}
这样的代码是错误的,编译器会提示cannot convert p (type *int) to type *int64
指针的强制类型转换需要用到unsafe包中的函数实现

```

```
package main

import "unsafe"
import "fmt"

func main() {
    var a int =10
    var b *int =&a
    var c *int64 = (*int64)(unsafe.Pointer(b))
    fmt.Println(*c)
}
```

golang中还有一中类型判断,类型断言

```
package main

import "fmt"

func main() {
    var a interface{} =10
    switch a.(type){
    case int:
            fmt.Println("int")
    case float32:
            fmt.Println("string")
    }
}
```

```
package main

import "fmt"

var b interface{} = "a"

func main() {
	fmt.Println("Hello, 世界")
 
	value, ok := b.(string)
if !ok {
    fmt.Println("It's not ok for type string")
    return
}
fmt.Println("The value is ", value)
}
```



### 0 字符串转换为数字

```
strconv.Atoi (表示ascii to integer)是把字符串转换成整型数的一个函数
```



### 1 **数字转化成字符串**

```go
var i = 65;
// 字符串的值不能被更改，但可以被替换
// 出错 变成 字符 "A" 了
str0 := string(i)

// 通过Itoa方法转换
str1 := strconv.Itoa(i)
 
// 通过Sprintf方法转换
str2 := fmt.Sprintf("%d", i)
fmt.Println(str0, str1, str2)

// A 65 65
```

### 2 字符串转换为byte 数组

```go
// string to ascii byte array
str := "ABC"
byteArray := []byte(str)
fmt.Println(byteArray);// [65 66 67]

// byte 数组转换为字符串 两种都可以
// s := string(byteArray[:])
s := fmt.Sprintf("%s", byteArray)
fmt.Println(s) // ABC
```

### 3 rune

rune在golang中是int32的别名，在各个方面都与int32相同。 被用来区分字符值和整数值。

```go
s:="hello你好"
fmt.Println(len(s))//输出长度为11
fmt.Println(len([]rune(s)))//输出长度为7
fmt.Println(([]rune(s)))// [104 101 108 108 111 20320 22909]
fmt.Println(len([]byte(s)))//输出长度为11
fmt.Println(([]byte(s)))// [104 101 108 108 111 228 189 160 229 165 189]

```

### 4 字符串分割和合并

```go
// 1 分割字符串
str1 := "1, 2, 3";
s1 := strings.Split(str1, ",");
fmt.Printf("\n%v %T", s1, s1);
// 2  合并为字符串   s1 [][]byte 
str2 := strings.Join(s1, ",")
fmt.Printf("\n%v", str2)
```



### 5 字符串判断

```go
// s 中是否以 prefix 开始
// strings.HasPrefix(s, prefix string) bool
// strings.HasSuffix(s, suffix string) bool
```

### **6 字符串查找的相关接口**

```go
// 判断给定字符串s中是否包含子串substr, 找到返回true, 找不到返回false
func Contains(s, substr string) bool

// 在字符串s中查找sep所在的位置, 返回位置值, 找不到返回-1
func Index(s, sep string) int

// 统计给定子串sep的出现次数, sep为空时, 返回1 + 字符串的长度
func Count(s, sep string) int
```

### 7 字符或子串在字符串中出现的位置

```go
// 在 s 中查找 sep 的第一次出现，返回第一次出现的索引
func Index(s, sep string) int
// chars中任何一个Unicode代码点在s中首次出现的位置
func IndexAny(s, chars string) int
// 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
func IndexFunc(s string, f func(rune) bool) int
// Unicode 代码点 r 在 s 中第一次出现的位置
func IndexRune(s string, r rune) int

// 有三个对应的查找最后一次出现的位置
func LastIndex(s, sep string) int
func LastIndexAny(s, chars string) int
func LastIndexFunc(s string, f func(rune) bool) int
```

### 8 时间

```go
//获取当前时间
   t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001
   fmt.Println(t)
 
   //获取当前时间戳
   fmt.Println(t.Unix()) //1531293019
 
   //获得当前的时间
   fmt.Println(t.Uninx().Format("2006-01-02 15:04:05"))  //2018-7-15 15:23:00
 
   //时间 to 时间戳
   loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
   tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
   fmt.Println(tt.Unix())                             //1531292871
 
   //时间戳 to 时间
   tm := time.Unix(1531293019, 0)
   fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19
 
   //获取当前年月日,时分秒
   y := t.Year()                 //年
   m := t.Month()                //月
   d := t.Day()                  //日
   h := t.Hour()                 //小时
   i := t.Minute()               //分钟
   s := t.Second()               //秒
   fmt.Println(y, m, d, h, i, s)
```

### 9 单词首字母

第一个单词首字母变大写：Ucfirst()，第一个单词首字母变小写：Lcfirst()

```go
import (
     "unicode"
)
 
func Ucfirst(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}
 
func Lcfirst(str string) string {
    for i, v := range str {
        return string(unicode.ToLower(v)) + str[i+1:]
    }
    return ""
}
```

### 10 浮点数最大最小值

Golang为什么没有整型的max/min方法

作为有一些经验的Golang开发者，你可能意识到了Golang并没有max/min方法来返回给定的两个或多个整型数值中的最大值或最小值。其他语言通常会在核心库中提供这类方法。 你有没有想过为什么Golang没有这么做？
Golang确实在`math`包中提供了max/min方法，但是仅用于对比float64类型。方法的签名如下：



```go
math.Min(float64, float64) float64
math.Max(float64, float64) float64
```

Golang为float64提供max/min方法是浮点类型的比较对于大部分开发者来说比较困难。由于涉及精度问题，浮点数的对比往往没有那么直接。所以Golang在`math`包中提供了用于浮点数对比的内建方法。
对于int/int64数据类型来说，max/min方法的实现就相当简单了，任何有基础编程经验的开发者都可以轻松的实现这两个方法：



```go
func Min(x, y int64) int64 {
 if x < y {
   return x
 }
 return y
}

func Max(x, y int64) int64 {
 if x > y {
   return x
 }
 return y
}
```

另外，为了尽可能保持Golang简洁干净，Golang并不支持泛型。所以既然已经有了对比float64数据类型的max/min方法，在同一个`math`包中就无法支持同名的但是参数类型不同的方法。也就是说下面的方法无法存在同一个包中。

### 11 高阶函数

```go
package main
 
import (
  "fmt"
  "reflect"
  "runtime"
  "time"
)
 
type SumFunc func(int64, int64) int64
 
func getFunctionName(i interface{}) string {
  return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
 
func timedSumFunc(f SumFunc) SumFunc {
  return func(start, end int64) int64 {
 
    defer func(t time.Time) {
      fmt.Printf("--- Time Elapsed (%s): %v ---\n", 
          getFunctionName(f), time.Since(t))
    }(time.Now())
 
    return f(start, end)
  }
}
 
func Sum1(start, end int64) int64 {
  var sum int64
  sum = 0
  if start > end {
    start, end = end, start
  }
  for i := start; i <= end; i++ {
    sum += i
  }
  return sum
}
 
func Sum2(start, end int64) int64 {
  if start > end {
    start, end = end, start
  }
  return (end - start + 1) * (end + start) / 2
}
 
func main() {
 
  sum1 := timedSumFunc(Sum1)
  sum2 := timedSumFunc(Sum2)
 
  fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}
```

