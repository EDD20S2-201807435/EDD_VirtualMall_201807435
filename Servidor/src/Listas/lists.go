package Listas

import (
	"fmt"
	"strconv"
)

type Node struct {
	Name string
	Age  int
	Next *Node
	Last *Node
}
type List struct {
	frist, last *Node
}

func NewList() *List {
	return &List{nil, nil}
}

func (this *List) Add(new *Node) {
	if this.frist == nil {
		this.frist = new
		this.last = new
	} else {
		this.last.Next = new
		new.Last = this.last
		this.last = new
	}
}

func (this *Node) To_string() string {
	return "Name: " + this.Name + "Edad: " + strconv.Itoa(this.Age)
}
func (this *List) To_string() string {
	var char string
	aux := this.frist
	for aux != nil {
		char += aux.To_string() + "\n"
		aux = aux.Next
	}
	return char
}

func (this *List) Print() {
	fmt.Println("Lista--------------")
	fmt.Println(this.To_string())
}
