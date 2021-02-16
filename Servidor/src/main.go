package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./Listas"
	"github.com/gorilla/mux"
)


func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my APP to Virtual Mall")
	fmt.Println("Lista")

}

type Message struct{
	Name string 'json:Nombre'
	Age int 'json:Edad'
}

func Add1(w http.ResponseWriter, r *http.Request){
	var ms Listas.Node
	reqBody,err:= ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"Error al Insertar")

	}
	json.Unmarshal(reqBody,&ms)
	fmt.Fprintln(w, ms.To_string())
}

func main() {
	lists := Listas.NewList()
	a := Listas.Node{"Milton", 21, nil, nil}
	b := Listas.Node{"Josue", 19, nil, nil}
	c := Listas.Node{"Rodriguez", 20, nil, nil}
	d := Listas.Node{"Valdez", 22, nil, nil}
	lists.Add(&a)
	lists.Add(&b)
	lists.Add(&c)
	lists.Add(&d)
	lists.Print()

	router := mux.NewRouter()
	router.HandleFunc("/", start).Methods("GET")
	router.HandleFunc("/agregar",Add1).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))


}
