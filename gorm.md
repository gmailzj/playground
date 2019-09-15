# gorm

## 代码生成器 gen

https://github.com/smallnest/gen

https://www.codercto.com/a/43809.html

gen工具从给定的 [数据库](http://www.codercto.com/category/database.html) 生成golang结构，以便在.go文件中使用。它支持 [gorm](https://github.com/jinzhu/gorm) 标签并实现一些可用的方法。它还可以为这些结构生成RESTful api。

通过从数据库中读取有关列结构的详细信息，gen生成具有所需列名，数据类型和注释的 [go](http://www.codercto.com/category/go.html) 兼容结构类型。

生成的数据类型包括对可空列 [sql.NullX类型](https://golang.org/pkg/database/sql/#NullBool) 或 [guregu null.X类型](https://github.com/guregu/null) 以及预期的基本内置go类型的支持。

gen基于Seth Shelnutt的 [db2struct](https://github.com/Shelnutt2/db2struct) 的工作而受到启发，而Db2Struct的基础/灵感来自ChimeraCoder的gojson包[gojson的工作](https://github.com/ChimeraCoder/gojson) 。

#### 获取

```
go get github.com/smallnest/gen
```

#### 使用

```
$ gen --connstr "root@tcp(127.0.0.1:3306)/employees?&parseTime=True" --database employees --json --gorm --guregu --rest
```

gen --connstr "root@tcp(127.0.0.1:3306)/fpmerchant?&parseTime=True" ~~--database fpmerchant~~ --json --gorm --guregu -t m_sell_client_new



## GORM

### GORM不定参数

// Topic 话题

```go
type Topic struct {
	gorm.Model
	Title         string    gorm:"index"
	Content       string    gorm:"type:text"
	ViewCount     int       json:"view_count"
	ReplyCount    int       json:"reply_count"
	UserID        int       gorm:"index" json:"user_id"
}

// QueryByUserID 根据UserID获取话题
func (t *Topic) QueryByUserID(userID int) (topics []Topic, err error) {
	db := config.DB.Where("user_id = ?", userID).Find(&topics).Error
	return
}

```

为了根据用户ID（UserID）获取到相应的话题（Topic），我们定义了一个QueryByUserID 的方法，为了表意这里刻意传入了一个参数；好像这个函数已经能帮我们解决model层与controller的分离。但是如果真正使用时会发现，其实这个函数有很大的局限性：它只能获取某个UserID的所有的话题，却无法实现对这个UserID下话题的进一步筛选（比如获取最近七天发布的话题）。

那么应该如何定义这个Query函数，使得我们使用时既优雅又功能强大呢？答案便是不定参数。

试想，我们把QueryByUserID函数以如下的方式定义：

```go
// QueryByUserID 根据UserID获取话题
func (t *Topic) QueryByUserID(userID int, args ...interface{}) (topics []Topic, err error) {
	db := config.DB.Where("user_id = ?", userID)
	if len(args) >= 2 {
		db = db.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		db = db.Where(args[0])
	}
	err = db.Find(&topics).Error
	return
}
```

我们这里在QueryByUserID的最后强行添加了一个args不定参数，从而接受调用时传入的附加的参数。同时在函数体内，我们根据args的长度对它进行不同方式的应用，这就非常优雅地扩展了QueryByUserID的功能。根据GORM的特性，大量使用了Go的反射特性，查看源码后知道这种用法完全可行。

在添加了不定参数的情况下，再调用QueryByUserID的时候，想要进一步精细地搜索话题，就方便多了。比如，想要获取UserID为1且最近一周发布的话题，可以这么写了：

```go
topics, err := (&Topic{}).QueryByUserID(1, "created_at > ?", time.Now().Add(-724time.Hour)
// ... 其他对topics 与 err 的处理
```

## 查询

```go
// 获取第一条记录，按主键排序
db.First(&user)
//// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取最后一条记录，按主键排序
db.Last(&user)
//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 获取所有记录
db.Find(&users)
//// SELECT * FROM users;

// 使用主键获取记录
db.First(&user, 10)
//// SELECT * FROM users WHERE id = 10;
```

### Where查询条件 (简单SQL)

```go
// 获取第一个匹配记录
db.Where("name = ?", "jinzhu").First(&user)
//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

// 获取所有匹配记录
db.Where("name = ?", "jinzhu").Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu';

db.Where("name <> ?", "jinzhu").Find(&users)

// IN
db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)

db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
```

### Where查询条件 (Struct & Map)

注意：当使用struct查询时，GORM将只查询那些具有值的字段

```go
// Struct
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// 主键的Slice
db.Where([]int64{20, 21, 22}).Find(&users)
//// SELECT * FROM users WHERE id IN (20, 21, 22);
```



### 错误处理

执行任何操作后，如果发生任何错误，GORM将其设置为`*DB`的`Error`字段

```go
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
    // 错误处理...
}

// 如果有多个错误发生，用`GetErrors`获取所有的错误，它返回`[]error`
db.First(&user).Limit(10).Find(&users).GetErrors()

// 检查是否返回RecordNotFound错误
db.Where("name = ?", "hello world").First(&user).RecordNotFound()

if db.Model(&user).Related(&credit_card).RecordNotFound() {
    // 没有信用卡被发现处理...
}
```



# 不用gorm

使用原生sql操作

```
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
```

**Gendry**是一个用于辅助操作数据库的Go包。基于`go-sql-driver/mysql`

https://github.com/didi/gendry



# 错误和异常

在Go的惯用法中，返回值不是整型等常用返回值类型，而是用了一个 error(interface类型)。

```go
type interface error {
    Error() string
}
```

在Go标准库中，Go提供了两种创建一个实现了error interface的类型的变量实例的方法：errors.New和fmt.Errorf：

```go
errors.New("your first error code")
fmt.Errorf("error value is %d\n", errcode)
```

这两个方法实际上返回的是同一个实现了error interface的类型实例，这个unexported类型就是errorString。顾名思义，这个error type仅提供了一个string的context！

```go
//$GOROOT/srcerrors/errors.go

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

错误和异常从语言机制上面讲，就是error和panic的区别

regexp包中有两个函数Compile和MustCompile，它们的声明如下：

```go
func Compile(expr string) (*Regexp, error)
func MustCompile(str string) *Regexp
```

同样的功能，不同的设计：

1. Compile函数基于错误处理设计，将正则表达式编译成有效的可匹配格式，适用于用户输入场景。当用户输入的正则表达式不合法时，该函数会返回一个错误。
2. MustCompile函数基于异常处理设计，适用于硬编码场景。当调用者明确知道输入不会引起函数错误时，要求调用者检查这个错误是不必要和累赘的。我们应该假设函数的输入一直合法，当调用者输入了不应该出现的输入时，就触发panic异常。

于是我们得到一个启示：**什么情况下用错误表达，什么情况下用异常表达，就得有一套规则，否则很容易出现一切皆错误或一切皆异常的情况。**