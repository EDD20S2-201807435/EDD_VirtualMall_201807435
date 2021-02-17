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

type Node_Tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Next         *Node_Tienda
	Last         *Node_Tienda
}

type Node_Departamento struct {
	Nombre  string
	Tiendas *Node_Tienda
}

type Node_Datos struct {
	Indice        string
	Departamentos [10]Node_Departamento
}

type Node_Dato struct {
	Datos [5]Node_Datos
}

type List struct {
	frist, last *Node
}

type List_Tienda struct {
	frist, last *Node_Tienda
}

type Vector_Departamnetos struct {
	
}

type Vector_Datos struct {
	vector_dato [5]Node_Datos
}
func NewList_vector() *ListVector{
	return vector_dep [100][4]Node_Departamento
}
func NewList() *List {
	return &List{nil, nil}
}
func NewLista_Tienda() *List_Tienda {
	return &List_Tienda{nil, nil}
}

func (this *List_Tienda) Add_Tienda(new *Node_Tienda) Node_Tienda {
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

var fila = 1

func (this Vector_Departamnetos) Add_Departementos(new Node_Departamento) {
	for i := 0; i < 10; i++ {
		if this.vector_dep[fila][i].Nombre == "" {
			this.vector_dep[fila][i] = new
		}
	}

}

func (this *Vector_Datos) Add_Datos(nodo Vector_Departamnetos) {

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
func (this *Vector_Departamnetos) Print_Dep() {
	fmt.Println("Lista--------------")

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(this.vector_dep[i][j].Nombre + " ")
			if this.vector_dep[i][j].Tiendas != nil {
				aux := this.vector_dep[i][j].Tiendas
				for aux != nil {
					fmt.Println("Nombre: " + aux.Nombre)
					fmt.Println("Descripcion " + aux.Descripcion)
					fmt.Println("Contacto " + aux.Contacto)
					fmt.Println("Calificacion " + strconv.Itoa(aux.Calificacion))
					aux = aux.Next
				}
			}

		}
	}
}
