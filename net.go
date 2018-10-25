package main

import (
    "strconv"
    "io"
    "net/http"
    "time"
)

// HelloServer the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    timestamp := time.Now().Unix()
    io.WriteString(w, strconv.FormatInt(timestamp, 10))
}

func main() {
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":2018", nil)
}