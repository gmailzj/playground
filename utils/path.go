package main

import (
    "path/filepath"
    "os"
    "os/exec"
    "strings"
    "log"
)

func GetAppPath() string {
    file, _ := exec.LookPath(os.Args[0])
    log.Println(file)
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))

    return path[:index]
}

func main() {
    log.Println(GetAppPath())
}