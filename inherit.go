package main

import (
    "fmt"
)

type People struct{}

type People2 struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

func (p *People) ShowC() {
    fmt.Println("showC")
}

func (p *People) ShowD() {
    fmt.Println("People:showD")
}

func (p *People2) ShowD() {
    fmt.Println("People2:showD")
}

type Teacher struct {
    People  //组合People
    People2 //组合People2
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}
func (t *Teacher) ShowC(arg string) {
    fmt.Println(arg)
}

func main() {
    t := Teacher{}

    //print showA
    //print showB
    t.ShowA()

    //print teacher showB
    t.ShowB()

    //print showB
    t.People.ShowB()

    //print test
    t.ShowC("test")

    //print showC
    t.People.ShowC()

    //因为组合方法中多次包含ShowD，所以调用时必须显示指定匿名方法
    //print People2:showD
    t.People2.ShowD()
}