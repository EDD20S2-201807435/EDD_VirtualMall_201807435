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
	Nombre  string     `json:Nombre`
	Tiendas [10]Tienda `json:Tiendas`
}

type Datos struct {
	Indice        string           `json:Indice`
	Departamentos [10]Departamento `json:Departamentos`
}

type Dato struct {
	Datos [5]Datos `json:Datos`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var ms Message
	var ms1 Datos
	var ms2 Departamento
	var ms3 Tienda
	var ms4 Dato
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	json.Unmarshal(reqBody, &ms4)
	json.Unmarshal(reqBody, &ms1)
	json.Unmarshal(reqBody, &ms2)
	json.Unmarshal(reqBody, &ms3)
	json.NewEncoder(w).Encode(ms)
	json.NewEncoder(w).Encode(ms4)
	json.NewEncoder(w).Encode(ms1)
	json.NewEncoder(w).Encode(ms2)
	json.NewEncoder(w).Encode(ms3)
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
	lists_tienda := Listas.NewLista_Tienda()
	a := Listas.Node{"Milton", 21, nil, nil}
	b := Listas.Node_Tienda{"Eskala", "Centro Comercial", "45875432", 3, nil, nil}
	lists.Add(&a)
	lists_tienda.Add_Tienda(&b)
	lists.Print()
	lists.Print_Dep()

	router := mux.NewRouter()
	router.HandleFunc("/", start).Methods("GET")
	router.HandleFunc("/cargartiendas", Add).Methods("POST")
	router.HandleFunc("/numero/{id}", number).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}
