package main
import(
	"net/http"
	"encoding/json"
	"./Listas"
	"fmt"
	"strconv"
	"io/ioutil"
	
)
type Node_Datos1_JSON struct {
	Datos[20] Node_Datos_JSON
}
type Node_Datos_JSON struct {
	Indice string
	Departamentos[50] Node_Departamento_JSON
}

type Node_Departamento_JSON struct {
	Nombre string
	Calificaciones[5] Node_Calificacion
}
type Node_Calificacion struct{
	Calificacion int
	Tiendas *Listas.List_Tienda
}

type data struct{
	Vector[2000] Listas.Calificacion
}
func CrearJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	Vector := Listas.Vector

	w.Header().Set("Content-Type","application/json")
	fmt.Println(strconv.Itoa(len(Vector)))
	
	
	
	var Vector1 [2000]Listas.Calificacion
		for i := 0; i < len(Vector); i++ {
			if Vector[i].Puntos != 0 {	
				Vector1[i] = Vector[i]
				
			}
		}
		dat := data{Vector1}
		json.NewEncoder(w).Encode(&dat)
		
		
	
}
var arbol = Listas.NewArbolB(5)
func InsertarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	reqBody, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	var prod Usu
	er := json.Unmarshal(reqBody, &prod)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	usuu :=make([]*Usuarios,len(prod.Usuarios))
	usuu = prod.Usuarios
	fmt.Println(arbol)
	for i := 0; i < len(usuu); i++ {
		fmt.Println(usuu[i].Nombre)
		arbol.InsertarB(Listas.NewKey(int64(usuu[i].Dpi), usuu[i].Nombre,usuu[i].Correo,usuu[i].Password,usuu[i].Cuenta))
	}
	arbol.GraficarB()
	
	
}



