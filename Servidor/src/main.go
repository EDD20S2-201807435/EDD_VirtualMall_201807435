package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"./Listas"
	"github.com/gorilla/mux"
)

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my APP to Virtual Mall")
	fmt.Println("Lista")

}

type Message struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type TE struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}
type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
}

type Departamento struct {
	Nombre  string    `json:"Nombre"`
	Tiendas [8]Tienda `json:"Tiendas"`
}

type Datos struct {
	Indice        string          `json:"Indice"`
	Departamentos [8]Departamento `json:"Departamentos"`
}

type Dato struct {
	Datos [5]Datos `json:"Datos"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	var dat Dato
	er := json.Unmarshal(reqBody, &dat)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	fmt.Println(dat.Datos[1].Indice)
	list_File := Listas.NewList_Datos()
	for i := 0; i < 5; i++ {
		//Si Hay un Nuevo Inidice
		if dat.Datos[i].Indice != "" {
			//Add Indices y Datos
			datoo := Listas.Node_Datos{dat.Datos[i].Indice, nil, nil}
			list_File.Add_Dato(&datoo)
			//Add Departamentos
			list_departamento := Listas.NewList_Departamentos()
			for j := 0; j < 8; j++ {
				if dat.Datos[i].Departamentos[j].Nombre != "" {
					//Departamento Existente
					depa := Listas.Node_Departamento{dat.Datos[i].Departamentos[j].Nombre, dat.Datos[i].Indice, nil, nil}
					list_departamento.Add_Departamento(&depa)
					//Add Calificaciones

					for k := 0; k < 5; k++ {
						cali := Listas.Calificacion{(k + 1), dat.Datos[i].Departamentos[j].Nombre, dat.Datos[i].Indice, nil}
						Listas.Add_Calificacion(cali, i)
					}
					//Add Tiendas a Calificaciones
					for w := 0; w < 5; w++ {
						listinda := Listas.NewLista_Tienda()
						for l := 0; l < 8; l++ {

							if dat.Datos[i].Departamentos[j].Tiendas[l].Nombre != "" {
								if dat.Datos[i].Departamentos[j].Tiendas[l].Calificacion == (w + 1) {
									store := Listas.Node_Tienda{dat.Datos[i].Departamentos[j].Tiendas[l].Nombre, dat.Datos[i].Departamentos[j].Tiendas[l].Descripcion, dat.Datos[i].Departamentos[j].Tiendas[l].Contacto, dat.Datos[i].Departamentos[j].Nombre, dat.Datos[i].Indice, dat.Datos[i].Departamentos[j].Tiendas[l].Calificacion, nil, nil}
									listinda.Add_Tienda(&store)
								}
							}
						}
						Listas.Search_Calificacion(dat.Datos[i].Departamentos[j].Nombre, dat.Datos[i].Indice, w+1, listinda.Return_Tienda())
					}

				}
			}

		}
	}

	Listas.Convertir_Matriz()

}
func Get_Arreglo(w http.ResponseWriter, r *http.Request) {
	Listas.Graficar()
}
func number(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, _ := strconv.Atoi(vars["id"])
	a := Message{"El numero que me mandaste es ", b}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(a)
}
func Tienda_Especifica(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	var Tienda TE
	er := json.Unmarshal(reqBody, &Tienda)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	fmt.Println(Tienda.Nombre)
	var tiendass *Listas.Node_Tienda
	tiendass = Listas.Tie_Esp(Tienda.Departamento, Tienda.Nombre, Tienda.Calificacion)
	if tiendass != nil {
		fmt.Fprintf(w, "Nombre: "+tiendass.Nombre)
		fmt.Fprintf(w, "Descripcion: "+tiendass.Descripcion)
		fmt.Fprintf(w, "Contacto: "+tiendass.Contacto)
		fmt.Fprintf(w, "Calificacion: "+strconv.Itoa(tiendass.Calificacion))

	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", start).Methods("GET")
	router.HandleFunc("/cargartiendas", Add).Methods("POST")
	router.HandleFunc("/getArreglo", Get_Arreglo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", Tienda_Especifica).Methods("POST")
	router.HandleFunc("/numero/{id}", number).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}
