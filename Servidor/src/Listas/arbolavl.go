package Listas

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Nodo struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
	Almacenammiento string
	Factor      int
	Izquierda   *Nodo
	Derecha     *Nodo
	
}

type Arbol struct {
	raiz *Nodo
}

func NewArbol() *Arbol {
	return &Arbol{nil}
}

func NewNodo(Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string,Almacenammiento string) *Nodo {
	return &Nodo{Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen,Almacenammiento, 0, nil, nil}
}

func rotacionII(n *Nodo, n1 *Nodo) *Nodo {
	n.Izquierda = n1.Derecha
	n1.Derecha = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func rotacionDD(n *Nodo, n1 *Nodo) *Nodo {
	n.Derecha = n1.Izquierda
	n1.Izquierda = n
	if n1.Factor == 1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func rotacionDI(n *Nodo, n1 *Nodo) *Nodo {
	n2 := n1.Izquierda
	n.Derecha = n2.Izquierda
	n2.Izquierda = n
	n1.Izquierda = n2.Derecha
	n2.Derecha = n1
	if n2.Factor == 1 {
		n.Factor = -1
	} else {
		n.Factor = 0
	}
	if n2.Factor == -1 {
		n1.Factor = 1
	} else {
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func rotacionID(n *Nodo, n1 *Nodo) *Nodo {
	n2 := n1.Derecha
	n.Izquierda = n2.Derecha
	n2.Derecha = n
	n1.Derecha = n2.Izquierda
	n2.Izquierda = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	} else {
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	} else {
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func insertar(raiz *Nodo, Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string,Almacenammiento string, hc *bool) *Nodo {
	var n1 *Nodo
	if raiz == nil {
		raiz = NewNodo(Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen,Almacenammiento)
		*hc = true
	} else if Codigo < raiz.Codigo {
		izq := insertar(raiz.Izquierda, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen,Almacenammiento, hc)
		raiz.Izquierda = izq
		if *hc {
			switch raiz.Factor {
			case 1:
				raiz.Factor = 0
				*hc = false
				break
			case 0:
				raiz.Factor = -1
				break
			case -1:
				n1 = raiz.Izquierda
				if n1.Factor == -1 {
					raiz = rotacionII(raiz, n1)
				} else {
					raiz = rotacionID(raiz, n1)
				}
				*hc = false
			}
		}
	} else if Codigo > raiz.Codigo {
		der := insertar(raiz.Derecha, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen,Almacenammiento, hc)
		raiz.Derecha = der
		if *hc {
			switch raiz.Factor {
			case 1:
				n1 = raiz.Derecha
				if n1.Factor == 1 {
					raiz = rotacionDD(raiz, n1)
				} else {
					raiz = rotacionDI(raiz, n1)
				}
				*hc = false
				break
			case 0:
				raiz.Factor = 1
				break
			case -1:
				raiz.Factor = 0
				*hc = false
			}

		}
	}
	return raiz
}

func (this *Arbol) Insertar(Nombre string, Codigo int, Descripcion string, Precio float64, Cantidad int, Imagen string,Almacenammiento string) {
	b := false
	a := &b
	this.raiz = insertar(this.raiz, Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen,Almacenammiento, a)
}

func (this *Arbol) Generar(name string) string{
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if this.raiz != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|<f1>%v  "+this.raiz.Nombre+": %v|<f2>\",color=green,style =filled];\n", &(*this.raiz), this.raiz.Codigo, this.raiz.Cantidad)
		this.generar(&cadena, (this.raiz), this.raiz.Izquierda, true)
		this.generar(&cadena, this.raiz, this.raiz.Derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), name)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "./"+name+"/"+name+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(name+"/"+name+".png", cmd, os.FileMode(mode))	
	return name+"/"+name+".png"
}

func (this *Arbol) generar(cadena *strings.Builder, padre *Nodo, actual *Nodo, Izquierda bool) {
	if actual != nil {
		fmt.Fprintf(cadena, "node%p[label=\"<f0>|<f1>%v "+actual.Nombre+": %v|<f2>\",color=green,style =filled];\n", &(*actual), actual.Codigo, actual.Cantidad)
		if Izquierda {
			fmt.Fprintf(cadena, "node%p:f0->node%p:f1\n", &(*padre), &(*actual))
		} else {
			fmt.Fprintf(cadena, "node%p:f2->node%p:f1\n", &(*padre), &(*actual))
		}
		this.generar(cadena, actual, actual.Izquierda, true)
		this.generar(cadena, actual, actual.Derecha, false)
	}
}



func (this *Arbol) MostrarArbol(){
	
	
	
}

func (this *Arbol) mostrarArbol( padre *Nodo, actual *Nodo, Izquierda bool) {
	
}



func (this *Arbol) Buscar_Producto(codigo int,cantidad int)*Nodo {
	var cadena strings.Builder
	if this.raiz != nil {
		if this.raiz.Codigo == codigo{
			this.raiz.Cantidad = this.raiz.Cantidad + cantidad
			return this.raiz
		}
		a := this.buscar_Producto(&cadena, (this.raiz), this.raiz.Izquierda, true,codigo,cantidad)
		b:=this.buscar_Producto(&cadena, this.raiz, this.raiz.Derecha, false,codigo,cantidad)
		if a != nil {
			return a
		}else if b!=nil{
			return b
		}
	}
	
	return nil
}

func (this *Arbol) buscar_Producto(cadena *strings.Builder, padre *Nodo, actual *Nodo, Izquierda bool,codigo int,cantidad int)*Nodo {
	if actual != nil {
		if Izquierda {
			if actual.Codigo == codigo {
				actual.Cantidad = actual.Cantidad + cantidad
				return actual
			}
		} else {
			if actual.Codigo == codigo {
				actual.Cantidad = actual.Cantidad + cantidad
				return actual
			}
		}
		this.buscar_Producto(cadena, actual, actual.Izquierda, true,codigo,cantidad)
		this.buscar_Producto(cadena, actual, actual.Derecha, false,codigo,cantidad)
	}
	return nil
}