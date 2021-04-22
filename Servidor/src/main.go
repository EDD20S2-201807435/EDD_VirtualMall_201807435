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
type Enlaces struct{
	Nombre string `json:"Nombre"`
	Distancia int `json:"Distancia"`
}
type Nodos struct{
	Nombre string `json:"Nombre"`
	Enlaces []*Enlaces `json:"Enlaces"`
}
type Grafo struct {
	Nodos []*Nodos `json:"Nodos"`
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
	Almacenamiento string  `json:"Almacenamiento"`
}
type Usuarios struct {
	Dpi      int64  `json:"Dpi"`
	Nombre     string    `json:"Nombre"`
	Correo string  `json:"Correo"`
	Password      string `json:"Password"`
	Cuenta    string     `json:"Cuenta"`
}
type Usu struct {
	Usuarios []*Usuarios `json:"Usuarios"`
}
type Inventario struct{
	Tienda string `json:"Tienda"`
    Departamento string `json:"Departamento"`
    Calificacion int `json:"Calificacion"`
    Productos[200] Productos `json:"Productos"`
}
type Inventarios struct {
	Inventario [50000]Inventario `json:"Inventarios"`
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
							tiendass.Productos.Insertar(prod.Inventario[i].Productos[j].Nombre,prod.Inventario[i].Productos[j].Codigo,prod.Inventario[i].Productos[j].Descripcion,prod.Inventario[i].Productos[j].Precio,prod.Inventario[i].Productos[j].Cantidad,prod.Inventario[i].Productos[j].Imagen,prod.Inventario[i].Productos[j].Almacenamiento)
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
type Dato_Tienda struct {
	Datos [100]Tien
}
type Imagen_arbol struct {
	url string
}
type Dato_imagen struct {
	url string
	Datos [1]Imagen_arbol
}
type Tien struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Departamento string
	Letra        string
	Calificacion int
	Logo         string
}

type Dato_Producto struct {
	Datos [100]Produ
}
type Produ struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
}
func Tienda_Espe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	b, _ := vars["id1"]
	c, _ := vars["id2"]
	d, _ := strconv.Atoi (vars["id3"])
	
	Vectoree := Listas.Listado_tiendas(b,c,d)
	var envio_Tiendas [100]Tien
	for i := 0; i<100;i++{
			if Vectoree[i].Nombre != "" {
				nodeo := Tien{Vectoree[i].Nombre,Vectoree[i].Descripcion,Vectoree[i].Contacto,Vectoree[i].Departamento,Vectoree[i].Letra,Vectoree[i].Calificacion,Vectoree[i].Logo}
				envio_Tiendas[i] = nodeo
			}
		
		
	}
	dat := Dato_Tienda{envio_Tiendas}
	
	json.NewEncoder(w).Encode(dat)
	
}
func Arbol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	b, _ := vars["id1"]
	c, _ := vars["id2"]
	d, _ := strconv.Atoi (vars["id3"])
	e, _ := vars["id4"]
	var avl [1]Imagen_arbol
	Vectoree := Listas.Tienda(b,c,d,e)
	envio_Tiendas := ""
	url:=""
	if Vectoree != nil {
		
		if Vectoree.Productos != nil {
			fmt.Println("Existe Arbol")
			envio_Tiendas =Vectoree.Productos.Generar(Vectoree.Nombre)
			dat1 := Imagen_arbol{"C:/Users/milto/EDD_VirtualMall_201807435/Servidor/src/"+envio_Tiendas}
			url="C:/Users/milto/EDD_VirtualMall_201807435/Servidor/src/"+envio_Tiendas
			avl[0] = dat1
		}
	}
	ar := Dato_imagen{url,avl}
	json.NewEncoder(w).Encode(ar)
	
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
func handler(w http.ResponseWriter, req *http.Request) {
    // ...
	enableCors(&w)
    // ...	
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func main() {
	
	
	router := mux.NewRouter()
	router.HandleFunc("/", start).Methods("GET")
	router.HandleFunc("/cargartiendas", Add).Methods("POST")
	router.HandleFunc("/getArreglo", Get_Arreglo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", Tienda_Especifica).Methods("POST")
	router.HandleFunc("/listadotiendas/{id1}/{id2}/{id3}", Tienda_Espe).Methods("GET")
    router.HandleFunc("/a",Add_Producto).Methods("POST")
	router.HandleFunc("/gete", Add_Pedido).Methods("POST")
	router.HandleFunc("/nem", Ingresar).Methods("GET")
	router.HandleFunc("/Arbol/{id1}/{id2}/{id3}/{id4}", Arbol).Methods("GET")
	
	router.HandleFunc("/crear", CrearJson)
	router.HandleFunc("/newusuario", InsertarCliente).Methods("POST")
	router.HandleFunc("/grafo", Add_Grafo).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}
