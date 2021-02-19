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
	Name string `json:Name`
	Age  int    `json:Age`
}

type Tienda struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	contacto     string `json:Contacto`
	Calificacion int    `json:Calificacion`
}

type Departamento struct {
	Nombre string `json:Nombre`
}

type Datos struct {
	Indice string `json:Indice`
}

type Dato struct {
	Datos string `json:Datos`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var ms Listas.Node

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	c := Dato{}
	var ell = json.Unmarshal(reqBody, &c)
	if ell != nil {
		fmt.Fprintf(w, "Error al Insertar")
	}
	json.Unmarshal(reqBody, &ms)
	fmt.Fprintln(w, ms.To_string())
	fmt.Fprintln(w, &c)

}

func number(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, _ := strconv.Atoi(vars["id"])
	a := Message{"El numero que me mandaste es ", b}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(a)
}

func main() {
	lists := Listas.NewList()
	//lists_tienda := Listas.NewLista_Tienda()
	//lists_tienda1 := Listas.NewLista_Tienda()
	//b := Listas.Node_Tienda{"Eskala", "Centro Comercial", "45875432", "Deportes", "A", 3, nil, nil}
	//c := Listas.Node_Tienda{"Eskala1", "Centro Comercial", "45875432", "Deportes", "A", 4, nil, nil}
	//d := Listas.Node_Tienda{"Eskala2", "Centro Comercial", "45875432", "Deportes", "A", 1, nil, nil}
	//e := Listas.Node_Tienda{"Eskala3", "Centro Comercial", "45875432", "Deportes", "A", 2, nil, nil}

	a := Listas.Node{"Milton", 21, nil, nil}

	lists.Add(&a)
	//lists_tienda.Add_Tienda(&b)
	//lists_tienda.Add_Tienda(&d)
	//lists_tienda1.Add_Tienda(&c)
	//lists_tienda1.Add_Tienda(&e)
	lists.Print()
	//lists_tienda.Print_Tienda()
	//g := Listas.Calificacion{1, nil}
	//h := Listas.Calificacion{2, lists_tienda.Return_Tienda()}
	//i := Listas.Calificacion{3, nil}
	//j := Listas.Calificacion{4, lists_tienda1.Return_Tienda()}
	//k := Listas.Calificacion{5, nil}
	//Listas.Add_Calificacion(g, 1)
	//Listas.Add_Calificacion(h, 1)
	//Listas.Add_Calificacion(i, 1)
	//Listas.Add_Calificacion(j, 1)
	//Listas.Add_Calificacion(k, 1)
	//Listas.Print_Vector()

	router := mux.NewRouter()
	router.HandleFunc("/", start).Methods("GET")
	router.HandleFunc("/cargartiendas", Add).Methods("POST")
	router.HandleFunc("/numero/{id}", number).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
	var hola Tienda
	fmt.Println(hola.Nombre)

}
