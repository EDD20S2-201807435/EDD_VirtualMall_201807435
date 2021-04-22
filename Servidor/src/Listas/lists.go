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
	Logo         string
	Productos *Arbol
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
func Tienda(indice string, depaa string, Puntos int,Nombre string) *Node_Tienda{
	
	for i := 0; i < len(Vector); i++ {
		if Vector[i].Departamento == depaa && Vector[i].Indice == indice && Vector[i].Puntos == Puntos{
			
			if Vector[i].Listatienda != nil{
				
				listado := Vector[i].Listatienda
				aux := listado.frist
				
				for aux != nil {
					if aux.Nombre == Nombre{
						return aux
					}
					aux = aux.Next
				}
				
			}else{
				
			}
		}
	}
	return nil

}
func Listado_tiendas(indice string, depaa string, Puntos int) [100]Node_Tienda{
	var Vector_retorno [100]Node_Tienda
	for i := 0; i < len(Vector); i++ {
		if Vector[i].Departamento == depaa && Vector[i].Indice == indice && Vector[i].Puntos == Puntos{
			
			if Vector[i].Listatienda != nil{
				
				listado := Vector[i].Listatienda
				aux := listado.frist
				con := 0
				for aux != nil {
					fmt.Println(*aux)
					Vector_retorno[con] = *aux
					
					con = con + 1
					aux = aux.Next
				}
				
			}else{
				
			}
		}
	}
	return Vector_retorno

}
func Listado_productos(indice string, depaa string, Puntos int,nombreTienda string) [100]Node_Tienda{
	var Vector_retorno [100]Node_Tienda
	for i := 0; i < len(Vector); i++ {
		if Vector[i].Departamento == depaa && Vector[i].Indice == indice && Vector[i].Puntos == Puntos{
			
			if Vector[i].Listatienda != nil{
				
				listado := Vector[i].Listatienda
				aux := listado.frist
				con := 0
				for aux != nil {
					if aux.Nombre == nombreTienda{

					}
					Vector_retorno[con] = *aux
					
					con = con + 1
					aux = aux.Next
				}
				
			}else{
				
			}
		}
	}
	return Vector_retorno

}
type Calificacion struct {
	Puntos       int
	Departamento string
	Indice       string
	Listatienda  *List_Tienda
}

var Matriz1 [2000][2000]Calificacion

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
	
	for i := 0; i < len(Matriz1); i++ {
		if Matriz1[No_Indice][i].Puntos == 0 {
			Matriz1[No_Indice][i] = new
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
	
}

func (this *List_Tienda) Print_Tienda() {
	
}
func (this *List_Tienda) Return_Tienda() *List_Tienda {
	aux := this.frist
	if aux != nil {
		return this
	}
	return nil
}

func Print_Vector() {
	
	for i := 0; i < len(Matriz1); i++ {
		for j := 0; j < len(Matriz1[i]); j++ {
			if Matriz1[i][j].Puntos != 0 {
				if Matriz1[i][j].Listatienda != nil {
					
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

	for i := 0; i < len(Matriz1); i++ {
		for j := 0; j < len(Matriz1[i]); j++ {
			if Matriz1[i][j].Puntos != 0 {
				if (Matriz1[i][j].Departamento == depa) && (Matriz1[i][j].Indice == indi) && (Matriz1[i][j].Puntos == cal) {
					Matriz1[i][j].Listatienda = Tiendaa
				}
			}
		}

	}
}

var Vector [4000000]Calificacion

func Convertir_Matriz() {
	var Contador int
	Contador = 0
	for i := 0; i < len(Matriz1); i++ {
		for j := 0; j < len(Matriz1[i]); j++ {
			if Matriz1[i][j].Puntos != 0 {
				Vector[Contador] = Matriz1[i][j]
				Contador++
			}
		}

	}
}

func Graficar(conteo int, centros int) string {

	var cadena strings.Builder
	if centros < len(Vector) && Vector[centros].Puntos != 0 {
		fmt.Fprintf(&cadena, "digraph G{\n")
		fmt.Fprintf(&cadena, "node[shape=box];\n")
		fmt.Fprintf(&cadena, "rankdir=TB;\n")
		fmt.Fprintf(&cadena, "graph[splines=polyline]\n")

		graficar(centros, &cadena, centros, 0)
		fmt.Fprintf(&cadena, "}")
		var name = strconv.Itoa(conteo) + ""
		guardarArchivo(cadena.String(), name)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "./"+name+"/"+name+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(name+"/"+name+".png", cmd, os.FileMode(mode))
		Graficar(conteo+1, centros+5)
	}

	return cadena.String()

}

func graficar(count int, s *strings.Builder, actual int, repetidor int) {
	if Vector[count].Puntos != 0 && repetidor < 5 {
		fmt.Fprintf(s, "node%p[label=\"%v|%v|%v\",color=blue,style =filled];\n", &Vector[count], Vector[count].Puntos, Vector[count].Departamento, Vector[count].Indice)
		if Vector[actual].Puntos != 0 {
			if &Vector[actual] != &Vector[count] {

				fmt.Fprintf(s, "{rank=same;node%p;node%p}\n", &Vector[count], &Vector[actual])
			}

			//Print Listas
			if Vector[count].Listatienda != nil {

				var validar *Node_Tienda
				validar = Vector[count].Listatienda.frist
				fmt.Fprintf(s, "node%p[label=\"%v|%v\",color=green,style =filled];\n", validar, validar.Nombre, validar.Calificacion)
				fmt.Fprintf(s, "node%p->node%p;\n", &Vector[count], validar)
				fmt.Fprintf(s, "node%p->node%p;\n", validar, &Vector[count])

				var validar1 *Node_Tienda
				validar1 = validar
				validar = validar.Next
				if validar != nil {

					for validar != nil {
						fmt.Fprintf(s, "node%p[label=\"%v|%v\",color=green,style =filled];\n", validar, validar.Nombre, validar.Calificacion)
						fmt.Fprintf(s, "node%p->node%p;\n", validar1, validar)
						fmt.Fprintf(s, "node%p->node%p;\n", validar, validar1)

						validar1 = validar
						validar = validar.Next
					}
				}

				//graficar(anterior.siguiente, s, anterior)
			}
		}
		graficar(count+1, s, count, repetidor+1)
	}
}

func guardarArchivo(cadena string, name string) {
	
	
	if _, err55 := os.Stat("./"+name); err55 != nil { if os.IsNotExist(err55) {
		err34 := os.RemoveAll("./"+name)
		if err34 != nil {
			fmt.Printf("Error eliminando carpeta con contenido: %v\n", err34)
		  } else {
			fmt.Println("Eliminada correctamente")
		}
		 } else { // other error 
			} 
		}
		os.Mkdir(name, os.ModeDir)

	f, err := os.Create(name + "/" + name + ".dot")
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

func Tie_Esp(depa string, Nombre string, cal int) *Node_Tienda {

	for i := 0; i < len(Vector); i++ {
		if Vector[i].Departamento == depa && Vector[i].Puntos == cal {
			if Vector[i].Listatienda != nil {
				aux := Vector[i].Listatienda.frist
				for aux != nil {
					if aux.Nombre == Nombre {
						return aux
					}
					aux = aux.Next
				}

			}
		}
	}

	return nil
}

func Graf_Vector() {

	for i := 0; i < len(Vector); i++ {
		if Vector[i].Puntos != 0 {
			
		}
	}

}


func Graficar_Api(conteo int, centros int) string {

	var cadena strings.Builder
	if centros < len(Vector) && Vector[centros].Puntos != 0 {
		graficar(centros, &cadena, centros, 0)
		
		var name = strconv.Itoa(conteo) + ""
		guardarArchivo(cadena.String(), name)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "./"+name+"/"+name+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(name+"/"+name+".png", cmd, os.FileMode(mode))
		Graficar(conteo+1, centros+5)
	}

	return cadena.String()

}

func graficar_api(count int, s *strings.Builder, actual int, repetidor int) {
	if Vector[count].Puntos != 0 && repetidor < 5 {
		fmt.Fprintf(s, "node%p[label=\"%v|%v|%v\",color=blue,style =filled];\n", &Vector[count], Vector[count].Puntos, Vector[count].Departamento, Vector[count].Indice)
		if Vector[actual].Puntos != 0 {
			if &Vector[actual] != &Vector[count] {

				fmt.Fprintf(s, "{rank=same;node%p;node%p}\n", &Vector[count], &Vector[actual])
			}

			//Print Listas
			if Vector[count].Listatienda != nil {

				var validar *Node_Tienda
				validar = Vector[count].Listatienda.frist
				fmt.Fprintf(s, "node%p[label=\"%v|%v\",color=green,style =filled];\n", validar, validar.Nombre, validar.Calificacion)
				fmt.Fprintf(s, "node%p->node%p;\n", &Vector[count], validar)
				fmt.Fprintf(s, "node%p->node%p;\n", validar, &Vector[count])

				var validar1 *Node_Tienda
				validar1 = validar
				validar = validar.Next
				if validar != nil {

					for validar != nil {
						fmt.Fprintf(s, "node%p[label=\"%v|%v\",color=green,style =filled];\n", validar, validar.Nombre, validar.Calificacion)
						fmt.Fprintf(s, "node%p->node%p;\n", validar1, validar)
						fmt.Fprintf(s, "node%p->node%p;\n", validar, validar1)

						validar1 = validar
						validar = validar.Next
					}
				}

				//graficar(anterior.siguiente, s, anterior)
			}
		}
		graficar(count+1, s, count, repetidor+1)
	}
}
