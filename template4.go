package main 

import (
    "html/template" 

    "github.com/gin-gonic/gin" 
) 

var (
    partial1 = `{{define "elm1"}}<div>element1</div>{{end}}` 
    partial2 = `{{define "elm2"}}<div>element2</div>{{end}}` 
    body  = `{{template "elm1"}}{{template "elm2"}}` 
) 

func main() { 
    // Or use `ParseFiles` to parse tmpl files instead 
    t := template.Must(template.New("elements").Parse(body)) 

    app := gin.Default() 
    app.GET("/", func(c *gin.Context) { 
     c.HTML(200, "elements", nil) 
    }) 
    app.Run(":8000") 
}
// https://gohugo.io/templates/go-templates/
