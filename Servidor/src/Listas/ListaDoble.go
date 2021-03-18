package Listas

import (
	"fmt"
)
type Node_Producto struct {
	Producto *Nodo
	Next *Node_Producto
	Last *Node_Producto
}
type List_Producto struct {
	frist, last *Node_Producto
}
func NewList_Producto() *List_Producto {
	return &List_Producto{nil, nil}
}

func (this *List_Producto) Add_Producto(new *Node_Producto) Node_Producto {
	if this.frist == nil {
		this.frist = new
		this.last = new
	} else {
		this.last.Next = new
		new.Last = this.last
		this.last = new
	}
	return *this.frist
}
func (this *List_Producto) Return_Producto() *List_Producto {
	aux := this.frist
	if aux != nil {
		return this
	}
	return nil
}
type Node_Month struct {
	Mes string
	Matriz_Dispersa  *Matriz
	Next *Node_Month
	Last *Node_Month
}
type List_Month struct {
	frist, last *Node_Month
}
func NewList_Month() *List_Month {
	return &List_Month{nil, nil}
}

type Node_Year struct {
	Year string
	Months  *List_Month
	Next *Node_Year
	Last *Node_Year
}
type List_Year struct {
	frist, last *Node_Year
}
func NewList_Year() *List_Year {
	return &List_Year{nil, nil}
}
var Listado_Years = NewList_Year()
func  Add_Year(new *Node_Year) Node_Year {
	this := Listado_Years
	if this.frist == nil {
		this.frist = new
		this.last = new
		
	} else {
		this.last.Next = new
		new.Last = this.last
		this.last = new
	}
	return *this.frist
}

func (this *List_Month) Add_Month(new *Node_Month) Node_Month {
	if this.frist == nil {
		this.frist = new
		this.last = new
	} else {
		this.last.Next = new
		new.Last = this.last
		this.last = new
	}
	return *this.frist
}
func (this *Node_Year) To_string_Year() Node_Year {
	return *this
}
func Existe_Year(Year string) Node_Year {
	aux := Listado_Years.frist
	for aux != nil {
		tmp := aux.To_string_Year()
		if tmp.Year == Year{
			return tmp
		}
		aux = aux.Next
	}
	yy :=Node_Year{"",nil,nil,nil}
	return yy
}

func (this *Node_Month) To_string_Month() Node_Month {
	return *this
}
func (this List_Month) Existe_Month(Mes string) Node_Month {
	aux := this.frist
	for aux != nil {
		tmp := aux.To_string_Month()
		if tmp.Mes == Mes{
			return tmp
		}
		aux = aux.Next
	}
	return Node_Month{"",nil,nil,nil}
}

func Calendario() Node_Year {
	aux := Listado_Years.frist
	for aux != nil {
		tmp := aux.To_string_Year()
		fmt.Println(""+tmp.Year)

		//Meses
		this := tmp.Months
		aux4 := this.frist
		for aux4 != nil {
		tmp4 := aux4.To_string_Month()
		fmt.Println(""+tmp4.Mes)
		nombree:=tmp.Year+"-"+tmp4.Mes
		tmp4.Matriz_Dispersa.Grafo(nombree)
		aux4 = aux4.Next
		}
		//Fin Meses
		aux = aux.Next
	}
	yy :=Node_Year{"",nil,nil,nil}
	return yy
}