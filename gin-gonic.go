package main

import "github.com/gin-gonic/gin"

import (
    "fmt"
    _ "encoding/json"
)

func main() {
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
	    id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "pong",
			"id": id,
		})
		fmt.Println(c.Request.Method)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Print(r.Routes())
	
}