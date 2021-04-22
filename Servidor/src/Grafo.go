package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)
type EnlacesGrafo struct{
	Nombre string
	Next *EnlacesGrafo
	Last *EnlacesGrafo
}
type EnlacesNodos struct{
	Nodo_anterior *EnlacesGrafo
	Nodo_siguiente *EnlacesGrafo
	Distancia int
	Next *EnlacesNodos
	Last *EnlacesNodos
}
type List_grafo struct {
	frist, last *EnlacesGrafo
}
type List_nodos_grafo struct {
	frist, last *EnlacesNodos
}
func NewList_grafo() *List_grafo {
	return &List_grafo{nil, nil}
}
func NewList_nodos_grafo() *List_nodos_grafo {
	return &List_nodos_grafo{nil, nil}
}
func (this *List_grafo) Add_grafo(new *EnlacesGrafo) EnlacesGrafo {
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
func (this *List_nodos_grafo) Add_enlaces(new *EnlacesNodos)EnlacesNodos {
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

func (this *List_grafo) Buscar_Nodos(nombre string)*EnlacesGrafo{
	
	aux := this.frist
	for aux != nil {
		if aux.Nombre == nombre {
			return aux
		}
		aux = aux.Next
	}
	return nil
	
}
func (this *List_grafo) To_string_Grafo(s *strings.Builder){
	
	aux := this.frist
	for aux != nil {
		fmt.Fprintf(s, "node%p[label=\"%v\",color=blue,style =filled];\n", &(*aux), aux.Nombre)
		aux = aux.Next
	}
	
}
func (this *List_nodos_grafo) To_string_Enlaces(s *strings.Builder){
	
	aux := this.frist
	for aux != nil {
		fmt.Fprintf(s, "node%p->node%p [label=\"%v\"];\n", &(*aux.Nodo_anterior), &(*aux.Nodo_siguiente),aux.Distancia)
		aux = aux.Next
	}
	
}

func Graficar_grafo(enlaces *List_nodos_grafo, grafo *List_grafo) string {

	var cadena strings.Builder
	
		fmt.Fprintf(&cadena, "digraph G{\n")
		fmt.Fprintf(&cadena, "layout = circo;\n")
		fmt.Fprintf(&cadena, "node [shape = circle,fontname = Helvetica];\n")
		
		
		grafo.To_string_Grafo(&cadena)
		enlaces.To_string_Enlaces(&cadena)
		fmt.Fprintf(&cadena, "}")
		var name = "Grafo"
		guardarArchivo1(cadena.String(), name)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "./"+name+"/"+name+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(name+"/"+name+".png", cmd, os.FileMode(mode))
		
	

	return cadena.String()

}
func guardarArchivo1(cadena string, name string) {
	
	
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