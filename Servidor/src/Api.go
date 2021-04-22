package main
import(
	"net/http"
	"encoding/json"
	"./Listas"
	"fmt"
	"strconv"
	"io/ioutil"
	"crypto/aes"
 "crypto/cipher"
 "encoding/hex"


	
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
		fmt.Fprintf(w, "")
	}
	usuu :=make([]*Usuarios,len(prod.Usuarios))
	usuu = prod.Usuarios
	fmt.Println(arbol)
	for i := 0; i < len(usuu); i++ {
		fmt.Println(usuu[i].Nombre)
		contra := usuu[i].Password
		arbol.InsertarB(Listas.NewKey(int64(usuu[i].Dpi), usuu[i].Nombre,usuu[i].Correo,contra,usuu[i].Cuenta))
	}
	arbol.GraficarB()
	
	
}
func encrypt(cadena string) string{
	key := []byte("keygopostmediumkeygopostmediumke")
	plaintext := []byte(cadena)
	block, err := aes.NewCipher(key)
	if err != nil {
	 panic(err.Error())
	}
	nonce := []byte("gopostmedium")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
	 panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Println("desde casa:"+string(ciphertext))
	return string(ciphertext)
   }
func decrypt(cadena string) string{
	key := []byte("keygopostmediumkeygopostmediumke")
	ciphertext, _ := hex.DecodeString(cadena)
	nonce := []byte("gopostmedium")
	block, err := aes.NewCipher(key)
	panic(err.Error())
	if err != nil {
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
	 panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
	 panic(err.Error())
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	return string(plaintext)
   }
func Add_Grafo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	reqBody, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	var prod Grafo
	er := json.Unmarshal(reqBody, &prod)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	usuu :=make([]*Nodos,len(prod.Nodos))
	usuu = prod.Nodos
	nuevografo := NewList_grafo()
	for i := 0; i < len(usuu); i++ {
		nodos := EnlacesGrafo{usuu[i].Nombre,nil,nil}
		nuevografo.Add_grafo(&nodos)
		
	}
	nuevoenlace := NewList_nodos_grafo()
	for i := 0; i < len(usuu); i++ {
		nod := nuevografo.Buscar_Nodos(usuu[i].Nombre)
		for j := 0; j < len(usuu[i].Enlaces); j++ {
			nod1 := nuevografo.Buscar_Nodos(usuu[i].Enlaces[j].Nombre)
			enlaces := EnlacesNodos{nod,nod1,usuu[i].Enlaces[j].Distancia,nil,nil}
			nuevoenlace.Add_enlaces(&enlaces)
		}
	
	}
	Graficar_grafo(nuevoenlace,nuevografo)
	
	
}


