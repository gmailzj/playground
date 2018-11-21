// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup. `defer` is often used where e.g.
// `ensure` and `finally` would be used in other languages.

package main

import "fmt"
import "os"

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
func main() {

    // Immediately after getting a file object with
    // `createFile`, we defer the closing of that file
    // with `closeFile`. This will be executed at the end
    // of the enclosing function (`main`), after
    // `writeFile` has finished.
    f := createFile("defer.txt")
    defer closeFile(f)
    writeFile(f)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    dir := "tmp/"
    exist, err := PathExists(dir)
    if err != nil {
        fmt.Printf("get dir error![%v]\n", err)
        panic(err)
    }
    fmt.Println(exist)
    if !exist {
       // 创建文件夹
        err := os.Mkdir(dir, os.ModePerm)
        if err != nil {
            panic(err)
        }
    }
    f, err := os.Create(dir + p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
