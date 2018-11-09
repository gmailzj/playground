package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func rot13(p byte) byte {
    switch {
    case p >= 'A' && p <= 'M':
        p += 13
    case p >= 'N' && p <= 'Z':
        p -= 13
    case p >= 'a' && p <= 'm':
        p += 13
    case p >= 'n' && p <= 'z':
        p -= 13
    }
    return p
}

func (m rot13Reader) Read(p []byte) (n int, err error) {
    ori := make([]byte, 50)
    i, err := m.r.Read(ori)  //读取写入的 io.Reader
    for index, value := range ori[:i] {
        p[index] = rot13(value)   //将写入的 io.Reader的值进行加密操作
    }
    return i, err
}

func main() {

    s := strings.NewReader("aaa!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
    io.Copy(os.Stdout, s) //没有值可以读取
    
    s = strings.NewReader("You cracked the code!")
    r = rot13Reader{s}
    io.Copy(os.Stdout, &r)
}