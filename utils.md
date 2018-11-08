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

// 两种都可以
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

