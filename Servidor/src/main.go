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
	Logo         string `json:"Logo"`
}

type Productos struct {
	Nombre      string  `json:"Nombre"`
	Codigo      int     `json:"Codigo"`
	Descripcion string  `json:"Descripcion"`
	Precio      float64 `json:"Precio"`
	Cantidad    int     `json:"Cantidad"`
	Imagen      string  `json:"Imagen"`
}
type Inventario struct{
	Tienda string `json:"Tienda"`
    Departamento string `json:"Departamento"`
    Calificacion int `json:"Calificacion"`
    Productos[200] Productos `json:"Productos"`
}
type Inventarios struct {
	Inventario [50000]Inventario `json:"Invetarios"`
}

type Departamento struct {
	Nombre  string      `json:"Nombre"`
	Tiendas [100]Tienda `json:"Tiendas"`
}

type Datos struct {
	Indice        string            `json:"Indice"`
	Departamentos [100]Departamento `json:"Departamentos"`
}

type Dato struct {
	Datos [100]Datos `json:"Datos"`
}
func Add_Producto(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	var prod Inventarios
	er := json.Unmarshal(reqBody, &prod)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	for i := 0; i < 50000; i++ {
		if prod.Inventario[i].Tienda != "" {
			var tiendass *Listas.Node_Tienda
			tiendass = Listas.Tie_Esp(prod.Inventario[i].Departamento, prod.Inventario[i].Tienda, prod.Inventario[i].Calificacion)
			if tiendass != nil {
				if tiendass.Productos == nil {
					tiendass.Productos = Listas.NewArbol()
				}
				for j := 0; j < 200; j++ {
					if prod.Inventario[i].Productos[j].Nombre != "" {
						comprabacion := tiendass.Productos.Buscar_Producto(prod.Inventario[i].Productos[j].Codigo,prod.Inventario[i].Productos[j].Cantidad)
						if comprabacion == nil {
							tiendass.Productos.Insertar(prod.Inventario[i].Productos[j].Nombre,prod.Inventario[i].Productos[j].Codigo,prod.Inventario[i].Productos[j].Descripcion,prod.Inventario[i].Productos[j].Precio,prod.Inventario[i].Productos[j].Cantidad,prod.Inventario[i].Productos[j].Imagen)
						}
						
					}
				}
				tiendass.Productos.Generar(prod.Inventario[i].Tienda)
			}
		}
	}
	


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
	for i := 0; i < 100; i++ {
		//Si Hay un Nuevo Inidice
		if dat.Datos[i].Indice != "" {
			//Add Indices y Datos
			datoo := Listas.Node_Datos{dat.Datos[i].Indice, nil, nil}
			list_File.Add_Dato(&datoo)
			//Add Departamentos
			list_departamento := Listas.NewList_Departamentos()
			for j := 0; j < 100; j++ {
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
						for l := 0; l < 100; l++ {

							if dat.Datos[i].Departamentos[j].Tiendas[l].Nombre != "" {
								if dat.Datos[i].Departamentos[j].Tiendas[l].Calificacion == (w + 1) {
									store := Listas.Node_Tienda{dat.Datos[i].Departamentos[j].Tiendas[l].Nombre, dat.Datos[i].Departamentos[j].Tiendas[l].Descripcion, dat.Datos[i].Departamentos[j].Tiendas[l].Contacto, dat.Datos[i].Departamentos[j].Nombre, dat.Datos[i].Indice, dat.Datos[i].Departamentos[j].Tiendas[l].Calificacion, dat.Datos[i].Departamentos[j].Tiendas[l].Logo,nil, nil, nil}
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
	Listas.Graf_Vector()
	Listas.Print_Vector()
}
func Get_Arreglo(w http.ResponseWriter, r *http.Request) {
	Listas.Graficar(0, 0)
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
	var Tienda1 TE
	er := json.Unmarshal(reqBody, &Tienda1)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	fmt.Println(Tienda1.Nombre)
	var tiendass *Listas.Node_Tienda
	tiendass = Listas.Tie_Esp(Tienda1.Departamento, Tienda1.Nombre, Tienda1.Calificacion)
	if tiendass != nil {
		var tt Tienda
		tt.Nombre = tiendass.Nombre
		tt.Descripcion = tiendass.Descripcion
		tt.Contacto = tiendass.Contacto
		tt.Calificacion = tiendass.Calificacion
		fmt.Fprintln(w, "Nombre: "+tiendass.Nombre)
		fmt.Fprintln(w, "Descripcion: "+tiendass.Descripcion)
		fmt.Fprintln(w, "Contacto: "+tiendass.Contacto)
		fmt.Fprintln(w, "Calificacion: "+strconv.Itoa(tiendass.Calificacion))
		w.Header().Set("Content-Type", "application/json")
		//json.Unmarshal(reqBody, tt)
		//json.NewEncoder(w).Enconde(&tt)

	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", start).Methods("GET")
	router.HandleFunc("/cargartiendas", Add).Methods("POST")
	router.HandleFunc("/getArreglo", Get_Arreglo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", Tienda_Especifica).Methods("POST")
	router.HandleFunc("/numero/{id}", number).Methods("GET")
    router.HandleFunc("/a",Add_Producto).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":3000", router))

}
