package main
import(
	"net/http"
	"encoding/json"
	"./Listas"
	"fmt"
	"strconv"
	
)
type Node_Datos1_JSON struct {
	Datos[3] Node_Datos_JSON
}
type Node_Datos_JSON struct {
	Indice string
	Departamentos[4] Node_Departamento_JSON
}

type Node_Departamento_JSON struct {
	Nombre string
	Calificaciones[5] Node_Calificacion
}
type Node_Calificacion struct{
	Calificacion int
	Tiendas *Listas.List_Tienda
}

func CrearJson(w http.ResponseWriter, r *http.Request) {
	Vector := Listas.Vector
	Contador := 0
	Contador_depa := 0
	Contador_datos := 0
	Indice := ""
	Departamento := ""
	var List_Cali [5]Node_Calificacion
	var List_Depa [4]Node_Departamento_JSON
	var List_Datos [3]Node_Datos_JSON
	fmt.Println(strconv.Itoa(len(Vector)))
	for i := 0; i < len(Vector); i++ {
		if Vector[i].Puntos != 0 {
			if Vector[i].Indice == Indice{
				
			if Contador < 4 {
				fmt.Println("Entra 5 "+strconv.Itoa(Contador))
				Indice = Vector[i].Indice
				Departamento = Vector[i].Departamento
				Cal := Node_Calificacion{Vector[i].Puntos,Vector[i].Listatienda}
				List_Cali[Contador]= Cal
				Contador++
			}else{
				Indice = Vector[i].Indice
				Departamento = Vector[i].Departamento
				Cal := Node_Calificacion{Vector[i].Puntos,Vector[i].Listatienda}
				List_Cali[Contador]= Cal
				
				fmt.Println("Crear un departamento")
				fmt.Println("%p",&List_Depa)
				De :=Node_Departamento_JSON{Departamento,List_Cali}
				List_Depa[Contador_depa] = De
				Contador_depa++
				i = i-1
				Contador = 0
			}
			}else{
				fmt.Println("Nuevo Indice")
				dato := Node_Datos_JSON{Indice,List_Depa}
				
				List_Datos[Contador_datos] = dato
				
				Indice = Vector[i].Indice
				fmt.Println(Indice)
				i = i-1
				
				
				Contador_depa = 0
				Contador_datos++
			}
		}
	}
	Datos := Node_Datos1_JSON{List_Datos}

	

	json.NewEncoder(w).Encode(Datos)
	
		
		
	
}



