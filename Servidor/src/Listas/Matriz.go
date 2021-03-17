package Listas

import (
	"fmt"
	"reflect"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type NodoPedido struct {
	ESTE, OESTE, NORTE, SUR interface{}
	Dia int
	Tienda string
	Departamento string
	Calificacion int
	Productos *Nodo
}

type NodoCabeceraVertical struct {
	ESTE  interface{}
	NORTE interface{}
	SUR   interface{}
	OESTE interface{}
	Departamento string
}

type NodoCabeceraHorizontal struct {
	ESTE, OESTE, SUR, NORTE interface{}
	Dia                   int
}

type Matriz struct {
	CabH *NodoCabeceraHorizontal
	CabV *NodoCabeceraVertical
}

func (this *Matriz) getHorizontal(dia int) interface{} {
	if this.CabH == nil {
		return nil
	}
	var aux interface{} = this.CabH
	for aux != nil {
		if aux.(*NodoCabeceraHorizontal).Dia == dia {
			return aux
		}
		aux = aux.(*NodoCabeceraHorizontal).ESTE
	}
	return nil
}

func (this *Matriz) getVertical(Departamento string) interface{} {
	if this.CabV == nil {
		return nil
	}
	var aux interface{} = this.CabV
	for aux != nil {
		if aux.(*NodoCabeceraVertical).Departamento == Departamento {
			return aux
		}
		aux = aux.(*NodoCabeceraVertical).SUR
	}
	return nil
}

func (this *Matriz) crearHorizontal(dia int) *NodoCabeceraHorizontal {
	if this.CabH == nil {
		nueva := &NodoCabeceraHorizontal{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Dia:  dia,
		}
		this.CabH = nueva
		return nueva
	}
	var aux interface{} = this.CabH
	if dia < aux.(*NodoCabeceraHorizontal).Dia {
		nueva := &NodoCabeceraHorizontal{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Dia:  dia,
		}
		nueva.ESTE = this.CabH
		this.CabH.OESTE = nueva
		this.CabH = nueva
		return nueva
	}
	for aux.(*NodoCabeceraHorizontal).ESTE != nil {
		if dia > aux.(*NodoCabeceraHorizontal).Dia && dia < aux.(*NodoCabeceraHorizontal).ESTE.(*NodoCabeceraHorizontal).Dia {
			nueva := &NodoCabeceraHorizontal{
				ESTE:  nil,
				OESTE: nil,
				SUR:   nil,
				NORTE: nil,
				Dia:  dia,
			}
			tmp := aux.(*NodoCabeceraHorizontal).ESTE
			tmp.(*NodoCabeceraHorizontal).OESTE = nueva
			nueva.ESTE = tmp
			aux.(*NodoCabeceraHorizontal).ESTE = nueva
			nueva.OESTE = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraHorizontal).ESTE
	}
	nueva := &NodoCabeceraHorizontal{
		ESTE:  nil,
		OESTE: nil,
		SUR:   nil,
		NORTE: nil,
		Dia:  dia,
	}
	aux.(*NodoCabeceraHorizontal).ESTE = nueva
	nueva.OESTE = aux
	return nueva
}

func (this *Matriz) crearVertical(departamento string) *NodoCabeceraVertical {
	if this.CabV == nil {
		nueva := &NodoCabeceraVertical{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Departamento: departamento,
		}
		this.CabV = nueva
		return nueva
	}
	var aux interface{} = this.CabV
	if departamento <= aux.(*NodoCabeceraVertical).Departamento {
		nueva := &NodoCabeceraVertical{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Departamento: departamento,
		}
		nueva.SUR = this.CabV
		this.CabV.NORTE = nueva
		this.CabV = nueva
		return nueva
	}
	for aux.(*NodoCabeceraVertical).SUR != nil {
		if departamento > aux.(*NodoCabeceraVertical).Departamento && departamento <= aux.(*NodoCabeceraVertical).SUR.(*NodoCabeceraVertical).Departamento {
			nueva := &NodoCabeceraVertical{
				ESTE:  nil,
				OESTE: nil,
				SUR:   nil,
				NORTE: nil,
				Departamento: departamento,
			}
			tmp := aux.(*NodoCabeceraVertical).SUR
			tmp.(*NodoCabeceraVertical).NORTE = nueva
			nueva.SUR = tmp
			aux.(*NodoCabeceraVertical).SUR = nueva
			nueva.NORTE = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraVertical).SUR
	}
	nueva := &NodoCabeceraVertical{
		ESTE:  nil,
		OESTE: nil,
		SUR:   nil,
		NORTE: nil,
		Departamento: departamento,
	}
	aux.(*NodoCabeceraVertical).SUR = nueva
	nueva.NORTE = aux
	return nueva
}

func (this *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, departamento string) interface{} {
	if cabecera.SUR == nil {
		return cabecera
	}
	aux := cabecera.SUR
	if departamento <= aux.(* NodoPedido).Departamento {
		return cabecera
	}
	for aux.(*NodoPedido).SUR != nil {
		if departamento > aux.(*NodoPedido).Departamento && departamento <= aux.(*NodoPedido).SUR.(*NodoPedido).Departamento {
			return aux
		}
		aux = aux.(*NodoPedido).SUR

	}
	if departamento <= aux.(*NodoPedido).Departamento {
		return aux.(*NodoPedido).NORTE
	}
	return aux
}
func (this *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dia int) interface{} {
	if cabecera.ESTE == nil {
		return cabecera
	}
	aux := cabecera.ESTE
	if dia <= aux.(*NodoPedido).Dia {
		return cabecera
	}
	for aux.(*NodoPedido).ESTE != nil {
		if dia > aux.(*NodoPedido).Dia && dia <= aux.(*NodoPedido).ESTE.(*NodoPedido).Dia {
			return aux
		}
		aux = aux.(*NodoPedido).ESTE
	}
	//tx:=aux.(*NodoPersona)
	if dia <= aux.(*NodoPedido).Dia {
		return aux.(*NodoPedido).OESTE
	}
	return aux
}

func (this *Matriz) Add(nueva *NodoPedido) {
	vertical := this.getVertical(nueva.Departamento)
	horizontal := this.getHorizontal(nueva.Dia)
	if vertical == nil {
		vertical = this.crearVertical(nueva.Departamento)
	}
	if horizontal == nil {
		horizontal = this.crearHorizontal(nueva.Dia)
	}
	izquierda := this.obtenerUltimoH(vertical.(*NodoCabeceraVertical), nueva.Dia)
	superior := this.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal), nueva.Departamento)
	if reflect.TypeOf(izquierda).String() == "*Listas.NodoPedido" {
		if izquierda.(*NodoPedido).ESTE == nil {
			izquierda.(*NodoPedido).ESTE = nueva
			nueva.OESTE = izquierda
		} else {
			tmp := izquierda.(*NodoPedido).ESTE
			izquierda.(*NodoPedido).ESTE = nueva
			nueva.OESTE = izquierda
			tmp.(*NodoPedido).OESTE = nueva
			nueva.ESTE = tmp
		}
	} else {
		if izquierda.(*NodoCabeceraVertical).ESTE == nil {
			izquierda.(*NodoCabeceraVertical).ESTE = nueva
			nueva.OESTE = izquierda
		} else {
			tmp := izquierda.(*NodoCabeceraVertical).ESTE
			izquierda.(*NodoCabeceraVertical).ESTE = nueva
			nueva.OESTE = izquierda
			tmp.(*NodoPedido).OESTE = nueva
			nueva.ESTE = tmp
		}
	}

	/*SUperior*/
	if reflect.TypeOf(superior).String() == "*Listas.NodoPedido" {
		if superior.(*NodoPedido).SUR == nil {
			superior.(*NodoPedido).SUR = nueva
			nueva.NORTE = superior
		} else {
			tmp := superior.(*NodoPedido).SUR
			superior.(*NodoPedido).SUR = nueva
			nueva.NORTE = superior
			tmp.(*NodoPedido).NORTE = nueva
			nueva.SUR = tmp
		}
	} else {
		if superior.(*NodoCabeceraHorizontal).SUR == nil {
			superior.(*NodoCabeceraHorizontal).SUR = nueva
			nueva.NORTE = superior
		} else {
			tmp := superior.(*NodoCabeceraHorizontal).SUR
			superior.(*NodoCabeceraHorizontal).SUR = nueva
			nueva.NORTE = superior
			tmp.(*NodoPedido).NORTE = nueva
			nueva.SUR = tmp
		}
	}
}

func (this *Matriz) Imprimir() {
	var aux interface{} = this.CabV
	for aux != nil {
		fmt.Print(aux.(*NodoCabeceraVertical).Departamento, "***************")
		tmp := aux.(*NodoCabeceraVertical).ESTE
		for tmp != nil {
			fmt.Printf("%v,%v------", tmp.(*NodoPedido).Dia, tmp.(*NodoPedido).Departamento)
			tmp = tmp.(*NodoPedido).ESTE
		}
		fmt.Print("\n")
		aux = aux.(*NodoCabeceraVertical).SUR
	}
}

func (this *Matriz) Imprimir2() {
	var aux interface{} = this.CabH
	for aux != nil {
		fmt.Print(aux.(*NodoCabeceraHorizontal).Dia, "*****************")
		tmp := aux.(*NodoCabeceraHorizontal).SUR
		for tmp != nil {
			fmt.Printf("%v,%v-------", tmp.(*NodoPedido).Dia, tmp.(*NodoPedido).Departamento)
			tmp = tmp.(*NodoPedido).SUR
		}
		fmt.Println("")
		aux = aux.(*NodoCabeceraHorizontal).ESTE
	}
}
func(this *Matriz)Grafo(){
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	fmt.Fprintf(&cadena, "rankdir=LR;\n")
	//
	var aux123 interface{} = this.CabH
	for aux123 != nil {
		fmt.Print(aux123.(*NodoCabeceraHorizontal).Dia, "*****************")
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|<f1>%v  "+strconv.Itoa(aux123.(*NodoCabeceraHorizontal).Dia)+": %v|<f2>\",color=green,style =filled];\n",  &(*aux123.(*NodoCabeceraHorizontal)), "", "")
		aux123 = aux123.(*NodoCabeceraHorizontal).ESTE
	}
	
		//Imprimir Cabeceras Horizontales 

	var aux interface{} = this.CabV
	for aux != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|<f1>%v  "+aux.(*NodoCabeceraVertical).Departamento+": %v|<f2>\",color=green,style =filled];\n",  &(*aux.(*NodoCabeceraVertical)), "", "")
		
		tmp := aux.(*NodoCabeceraVertical).ESTE
		tmp1 := aux.(*NodoCabeceraVertical).ESTE
		if tmp != nil {
			fmt.Fprintf(&cadena, "node%p[label=\"<f0>|<f1>%v  "+tmp.(*NodoPedido).Departamento+": %v|<f2>\",color=green,style =filled];\n",  &(*tmp.(*NodoPedido)), tmp.(*NodoPedido).Dia, "")
			fmt.Fprintf(&cadena, "node%p->node%p;\n",  &(*aux.(*NodoCabeceraVertical)), &(*tmp.(*NodoPedido)))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(*tmp.(*NodoPedido)), &(*aux.(*NodoCabeceraVertical)))
			tmp1 = tmp
			tmp = tmp.(*NodoPedido).ESTE
		}
		for tmp != nil {
			fmt.Printf("%v,%v------", tmp.(*NodoPedido).Dia, tmp.(*NodoPedido).Departamento)
			fmt.Fprintf(&cadena, "node%p[label=\"<f0>|<f1>%v  "+tmp.(*NodoPedido).Departamento+": %v|<f2>\",color=green,style =filled];\n",  &(*tmp.(*NodoPedido)), tmp.(*NodoPedido).Dia, "")
			fmt.Fprintf(&cadena, "node%p->node%p;\n",  &(*tmp1.(*NodoPedido)), &(*tmp.(*NodoPedido)))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(*tmp.(*NodoPedido)), &(*tmp1.(*NodoPedido)))
			fmt.Fprintf(&cadena, "{rank=same;node%p;node%p}\n", &(*tmp.(*NodoPedido)),&(*tmp1.(*NodoPedido)))
			tmp1 = tmp
			tmp = tmp.(*NodoPedido).ESTE
		}
		fmt.Print("\n")
		t:=aux
		aux = aux.(*NodoCabeceraVertical).SUR
		if aux != nil {
			fmt.Fprintf(&cadena, "{rank=same;rankdir=LR;node%p;node%p}\n", &(*aux.(*NodoCabeceraVertical)),&(*t.(*NodoCabeceraVertical)))
		}
		
	}
		
	//Imprimir Cabecera Vertical
	var aux12 interface{} = this.CabH
	for aux12 != nil {
		fmt.Print(aux12.(*NodoCabeceraHorizontal).Dia, "*****************")
		tmp12 := aux12.(*NodoCabeceraHorizontal).SUR
		tmp112 := aux12.(*NodoCabeceraHorizontal).SUR
		if tmp12 != nil {
			fmt.Fprintf(&cadena, "node%p->node%p;\n",  &(*aux12.(*NodoCabeceraHorizontal)), &(*tmp12.(*NodoPedido)))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(*tmp12.(*NodoPedido)), &(*aux12.(*NodoCabeceraHorizontal)))
			tmp112 = tmp12
			tmp12 = tmp12.(*NodoPedido).SUR
		}
		for tmp12 != nil {
			fmt.Printf("%v,%v-------", tmp12.(*NodoPedido).Dia, tmp12.(*NodoPedido).Departamento)
			fmt.Fprintf(&cadena, "node%p->node%p;\n",  &(*tmp112.(*NodoPedido)), &(*tmp12.(*NodoPedido)))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(*tmp12.(*NodoPedido)), &(*tmp112.(*NodoPedido)))
			fmt.Fprintf(&cadena, "{rank=same;rankdir=LR;node%p;node%p}\n", &(*tmp12.(*NodoPedido)),&(*tmp112.(*NodoPedido)))
			tmp112 = tmp12
			tmp12 = tmp12.(*NodoPedido).SUR
		}
		fmt.Println("")
		aux12 = aux12.(*NodoCabeceraHorizontal).ESTE
	}
	name := "MAtrizDispersa"
	fmt.Fprintf(&cadena, "}\n")
	guardarArchivo(cadena.String(), name)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "./"+name+"/"+name+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(name+"/"+name+".png", cmd, os.FileMode(mode))	
}