package Listas

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

//eliminar
type List struct {
	frist, last *Node
}

type List_Tienda struct {
	frist, last *Node_Tienda
}

type Calificacion struct {
	Puntos       int
	Departamento string
	Indice       string
	Listatienda  *List_Tienda
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
func NewList_Datos() *List_Datos {
	return &List_Datos{nil, nil}
}
func NewList_Departamentos() *List_Departamentos {
	return &List_Departamentos{nil, nil}
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

func (this *List_Datos) Add_Dato(new *Node_Datos) Node_Datos {
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
	aux := this.frist
	if aux != nil {
		return this
	}
	return nil
}

func Print_Vector() {
	fmt.Println("Lista--------------")
	for i := 0; i < 5; i++ {
		for j := 0; j < 30; j++ {
			if Matriz[i][j].Puntos != 0 {
				fmt.Println("Calificacion " + strconv.Itoa(Matriz[i][j].Puntos) + "Departamento: " + Matriz[i][j].Departamento)
				if Matriz[i][j].Listatienda != nil {
					fmt.Println(Matriz[i][j].Listatienda.To_string_Tienda())
				}
			}
		}

	}
}

func (this *Node_Datos) To_string_Dato() string {
	return "Indice: " + this.Indice
}
func (this *List_Datos) To_string_Dato() string {
	var char string
	aux := this.frist
	for aux != nil {
		char += aux.To_string_Dato() + "\n"
		aux = aux.Next
	}
	return char
}

func (this *List_Datos) Print_Dato() {

	fmt.Println(this.To_string_Dato())
}

func Search_Calificacion(depa string, indi string, cal int, Tiendaa *List_Tienda) {

	for i := 0; i < 5; i++ {
		for j := 0; j < 30; j++ {
			if Matriz[i][j].Puntos != 0 {
				if (Matriz[i][j].Departamento == depa) && (Matriz[i][j].Indice == indi) && (Matriz[i][j].Puntos == cal) {
					Matriz[i][j].Listatienda = Tiendaa
				}
			}
		}

	}
}

var Vector [100]Calificacion

func Convertir_Matriz() {
	var Contador int
	Contador = 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 30; j++ {
			if Matriz[i][j].Puntos != 0 {
				Vector[Contador] = Matriz[i][j]
				Contador++
			}
		}

	}
}

func Graficar() string {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=record];\n")
	graficar(0, &cadena, 0)
	fmt.Fprintf(&cadena, "}")
	guardarArchivo(cadena.String())
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./lista.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("outfile.png", cmd, os.FileMode(mode))

	return cadena.String()

}

func graficar(count int, s *strings.Builder, actual int) {
	if Vector[count].Puntos != 0 {
		fmt.Fprintf(s, "node%p[label=\"%v|%v\"];\n", &Vector[actual], Vector[actual].Puntos, Vector[actual].Departamento)
		if Vector[actual].Puntos != 0 {
			fmt.Fprintf(s, "node%p->node%p;\n", &Vector[actual], &Vector[count])
			fmt.Fprintf(s, "node%p->node%p;\n", &Vector[count], &Vector[actual])
			//Print Listas
			if Vector[count].Listatienda != nil {
				fmt.Fprintf(s, "node%p[label=\"%v|%v\"];\n", &Vector[count].Listatienda.frist, Vector[count].Listatienda.frist.Nombre, Vector[count].Listatienda.frist.Calificacion)
				fmt.Fprintf(s, "node%p->node%p;\n", &Vector[count], &Vector[count].Listatienda.frist)
				fmt.Fprintf(s, "node%p->node%p;\n", &Vector[count].Listatienda.frist, &Vector[count])
				graficar_Tiendas(Vector[count].Listatienda.frist.Next, s, Vector[count].Listatienda.frist)
				//graficar(anterior.siguiente, s, anterior)
			}
		}
		graficar(count+1, s, count)
	}
}

func graficar_Tiendas(anterior *Node_Tienda, s *strings.Builder, actual *Node_Tienda) {
	if anterior != nil {
		fmt.Fprintf(s, "node%p[label=\"%v|%v\"];\n", &(*anterior), anterior.Nombre, anterior.Calificacion)
		if actual != nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*actual), &(*anterior))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*anterior), &(*actual))
		}
		graficar_Tiendas(anterior.Next, s, anterior)
	}
}

func guardarArchivo(cadena string) {
	f, err := os.Create("lista.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
