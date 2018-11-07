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

