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
	Departamento string
	Letra        string
	Calificacion int
	Next         *Node_Tienda
	Last         *Node_Tienda
}

type Node_Departamento struct {
	Nombre string
	Indice string
	Next   *Node_Departamento
	Last   *Node_Departamento
}

type Node_Datos struct {
	Indice string
	Next   *Node_Datos
	Last   *Node_Datos
}

type List struct {
	frist, last *Node
}

type List_Tienda struct {
	frist, last *Node_Tienda
}

type Calificacion struct {
	Puntos      int
	Listatienda *List_Tienda
}

var Matriz [5][30]Calificacion

type List_Departamentos struct {
	frist, last *Node_Departamento
}

type List_Datos struct {
	frist, last *Node_Datos
}

func NewList_vector() *List_Tienda {
	return &List_Tienda{nil, nil}
}

//eliminar
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

func (this *List_Departamentos) Add_Departamento(new *Node_Departamento) Node_Departamento {
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

func (this *List_Datos) Add_Departamento(new *Node_Datos) Node_Datos {
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

func Add_Calificacion(new Calificacion, No_Indice int) Calificacion {
	for i := 0; i < 30; i++ {
		if Matriz[No_Indice][i].Puntos == 0 {
			Matriz[No_Indice][i] = new
			return new
		}
	}
	return new
}

var fila = 1

//eliminar
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

func (this *Node_Tienda) To_string_Tienda() string {
	return "Name: " + this.Nombre + "Edad: " + strconv.Itoa(this.Calificacion)
}
func (this *List_Tienda) To_string_Tienda() string {
	var char string
	aux := this.frist
	for aux != nil {
		char += aux.To_string_Tienda() + "\n"
		aux = aux.Next
	}
	return char
}

func (this *List) Print() {
	fmt.Println("Lista--------------")
	fmt.Println(this.To_string())
}

func (this *List_Tienda) Print_Tienda() {
	fmt.Println("Lista--------------")
	fmt.Println(this.To_string_Tienda())
}
func (this *List_Tienda) Return_Tienda() *List_Tienda {
	return this
}

func Print_Vector() {
	fmt.Println("Lista--------------")
	for i := 0; i < 25; i++ {
		if Matriz[1][i].Puntos != 0 {
			fmt.Println("Calificacion " + strconv.Itoa(Matriz[1][i].Puntos))
			if Matriz[1][i].Listatienda != nil {
				fmt.Println(Matriz[1][i].Listatienda.To_string_Tienda())
			}
		}

	}
}
