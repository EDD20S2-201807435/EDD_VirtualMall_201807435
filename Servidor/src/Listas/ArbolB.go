package Listas
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	
)
type ArbolB struct{
	k int
	Raiz *NodoB
}
func NewArbolB(nivel int)*ArbolB{
	a:= ArbolB{nivel,nil}
	nodoraiz := NewNodoB(nivel)
	a.Raiz = nodoraiz
	return &a
}
func (this *ArbolB)InsertarB(newKey *Key){
	if this.Raiz.Keys[0] == nil{
		this.Raiz.Colocar(0,newKey)
	}else if this.Raiz.Keys[0].Izquierdo == nil{
		lugarinsertado:=-1
		node:=this.Raiz
		lugarinsertado= this.colocarNodo(node,newKey)
		if lugarinsertado != -1 {
			if lugarinsertado == node.Max-1{
				middle := node.Max / 2
				llavecentral := node.Keys[middle]
				derecho := NewNodoB(this.k)
				izquierdo := NewNodoB(this.k)
				indiceizquierdo := 0
				indicederecho := 0
				for j := 0; j < node.Max; j++ {
					if node.Keys[j].Dpi < llavecentral.Dpi {
						izquierdo.Colocar(indiceizquierdo,node.Keys[j])
						indiceizquierdo++
						node.Colocar(j,nil)
					}else if node.Keys[j].Dpi > llavecentral.Dpi{
						derecho.Colocar(indicederecho,node.Keys[j])
						indicederecho++
						node.Colocar(j,nil)
					}
				}
				node.Colocar(middle,nil)
				this.Raiz = node
				this.Raiz.Colocar(0,llavecentral) 
				izquierdo.NodoPadre = this.Raiz
				derecho.NodoPadre = this.Raiz
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho 
			}
		}

	}else if this.Raiz.Keys[0].Izquierdo != nil {
		node := this.Raiz
		for node.Keys[0].Izquierdo!=nil {
			loop := 0
			for i := 0; i < node.Max; i,loop = i+1,loop+1 {
				if node.Keys[i] != nil{
					if node.Keys[i].Dpi > newKey.Dpi {
						node = node.Keys[i].Izquierdo
						break
					}
				}else{
					node = node.Keys[i-1].Derecho
					break
				}
			}
			if loop ==node.Max {
				node = node.Keys[loop-1].Derecho
			}
		}
		indiceColocado := this.colocarNodo(node,newKey)
		if indiceColocado == node.Max-1 {
			for node.NodoPadre != nil{
				indicemedio := node.Max/2
				llavecentral:=node.Keys[indicemedio]
				izquierdo:=NewNodoB(this.k)
				derecho:=NewNodoB(this.k)
				indiceizquierdo,indicederecho := 0,0
				for i := 0; i < node.Max; i++ {
					if node.Keys[i].Dpi < llavecentral.Dpi {
						izquierdo.Colocar(indiceizquierdo,node.Keys[i])
						indiceizquierdo++
						node.Colocar(i,nil)
					}else if node.Keys[i].Dpi > llavecentral.Dpi{
						derecho.Colocar(indicederecho,node.Keys[i])
						indicederecho++
						node.Colocar(i,nil)
					}
				}
				node.Colocar(indicemedio,nil)
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
				node = node.NodoPadre
				izquierdo.NodoPadre = node
				derecho.NodoPadre = node
				for i := 0; i < izquierdo.Max; i++ {
					if izquierdo.Keys[i] != nil {
						if izquierdo.Keys[i].Izquierdo !=nil {
							izquierdo.Keys[i].Izquierdo.NodoPadre = izquierdo
						}
						if izquierdo.Keys[i].Derecho != nil {
							izquierdo.Keys[i].Derecho.NodoPadre = izquierdo
						}
					}
					
				}
				for i := 0; i < derecho.Max; i++ {
					if derecho.Keys[i] != nil {
						if derecho.Keys[i].Izquierdo != nil {
							derecho.Keys[i].Izquierdo.NodoPadre=derecho 
						}
						if derecho.Keys[i].Derecho != nil {
							derecho.Keys[i].Derecho.NodoPadre = derecho
						}
					}
				}
				lugarcolocado:=this.colocarNodo(node,llavecentral)
				if lugarcolocado == node.Max-1 {
					if node.NodoPadre == nil{
						indicecentralraiz:=node.Max/2
						llavecentralraiz := node.Keys[indicecentralraiz]
						izquierdoraiz:=NewNodoB(this.k)
						derechoraiz:=NewNodoB(this.k)
						indicederechoraiz, indiceizquierdoraiz := 0,0
						for i := 0; i < node.Max; i++ {
							if node.Keys[i].Dpi < llavecentralraiz.Dpi {
								izquierdoraiz.Colocar(indiceizquierdoraiz,node.Keys[i])
								indiceizquierdoraiz++
								node.Colocar(i,nil)
							}else if node.Keys[i].Dpi > llavecentralraiz.Dpi{
								derechoraiz.Colocar(indicederechoraiz,node.Keys[i])
								indicederechoraiz++
								node.Colocar(i,nil)
							}
						}
						node.Colocar(indicecentralraiz,nil)
						node.Colocar(0,llavecentralraiz)
						for i := 0; i < this.k; i++ {
							if izquierdoraiz.Keys[i]!=nil {
								izquierdoraiz.Keys[i].Izquierdo.NodoPadre=izquierdoraiz
								izquierdoraiz.Keys[i].Derecho.NodoPadre=izquierdoraiz
							}
						}
						for i := 0; i < this.k; i++ {
							if derechoraiz.Keys[i]!=nil {
								derechoraiz.Keys[i].Izquierdo.NodoPadre=derechoraiz
								derechoraiz.Keys[i].Derecho.NodoPadre=derechoraiz
							}
						}
						llavecentralraiz.Izquierdo = izquierdoraiz
						llavecentralraiz.Derecho = derechoraiz
						izquierdoraiz.NodoPadre=node
						derechoraiz.NodoPadre = node
						this.Raiz = node
					}
					continue
				}else{
					break
				}
			}
		}
	}
}

func(this *ArbolB)colocarNodo(node *NodoB,newkey *Key)int {
	index := -1
	for i:= 0; i<node.Max;i++ {
		if node.Keys[i] == nil {
			placed := false
			for j := i-1; j >= 0; j-- {
				if node.Keys[j].Dpi > newkey.Dpi {
					node.Colocar(j+1,node.Keys[j])
				}else{
					node.Colocar(j+1,newkey)
					node.Keys[j].Derecho = newkey.Izquierdo
					if (j+2) < this.k && node.Keys[j+2] != nil{
						node.Keys[j+2].Izquierdo = newkey.Derecho
					}
					placed = true
					break
				}
			}
			if placed ==false {
				node.Colocar(0,newkey)
				node.Keys[1].Izquierdo=newkey.Derecho
			}
			index = i
			break
		}
	}
	return index
}
func(this *ArbolB)GraficarB(){
	builder := strings.Builder{}
	fmt.Fprintf(&builder,"digraph G{\nnode[shape=record]\n")
	m := make(map[string]*NodoB)
	graficarB(this.Raiz,&builder,m,nil,0)
	fmt.Fprintf(&builder,"}")
	guardarArchivoB(builder.String())
	generarimagen("arbol.png")
}
func graficarB(actual *NodoB,cad *strings.Builder,arr map[string]*NodoB,padre *NodoB,pos int){
	if actual == nil {
		return
	}
	j:= 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	}else{
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad,"node%p[label=\"",&(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] ==nil {
			return
		}else{
			if enlace {
				if i !=actual.Max-1 {
					fmt.Fprintf(cad,"<f%d>|",j)
				}else{
					fmt.Fprintf(cad,"<f%d>",j)
					break
				}
				enlace = false
				i--
				j++
			}else{
				fmt.Fprintf(cad,"<f%d>|{%d|"+actual.Keys[i].Cuenta+"|"+actual.Keys[i].Password+"}|",j,actual.Keys[i].Dpi)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad,"<f%d>",j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad,"\"]\n")
	ji:=0
	for i := 0; i < actual.Max-4; i++ {
		fmt.Println(actual.Max)
		if actual.Keys == nil {
			break
		}
		
		graficarB(actual.Keys[i].Izquierdo,cad,arr,actual,ji)
		ji++
		ji++
		graficarB(actual.Keys[i].Derecho,cad,arr,actual,ji)
		ji++
		ji--
		
		
	}
	if padre != nil {
		fmt.Fprintf(cad,"node%p:f%d->node%p\n",&(*padre),pos,&(*actual))
	}

}
func guardarArchivoB(cadena string){
	f, err :=os.Create("diagrama.dot")
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
	fmt.Println(l,"bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return	
	}
}
func generarimagen(nombre string){
	path, _ := exec.LookPath("dot")
	cmd, _ :=exec.Command(path, "-Tpng","./diagrama.dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre,cmd,os.FileMode(mode))
}