package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"./Listas"
)

type CodigoProductos struct{
	Codigo int `json:"Codigo"`
} 

type Pedido struct{
	Fecha string `json:"Fecha"`
    Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
    Calificacion int `json:"Calificacion"`
    CodigoProductos[200] CodigoProductos `json:"Productos"`
}
type Pedidos struct {
	Pedido [50000]Pedido `json:"Pedidos"`
}
func Ingresar(w http.ResponseWriter, r *http.Request){
	dispersa := &Listas.Matriz{nil, nil}
	listProducto := Listas.NewList_Producto()
	produc := &Listas.NodoPedido{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Dia: 11, Tienda: "Samsung", Departamento:"Tecnologia",Calificacion:5,Productos:listProducto}
	dispersa.Add(produc)
	
	dispersa.Imprimir()
	dispersa.Imprimir2()
	dispersa.Grafo("Prueba")
}
func Add_Pedido(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		fmt.Fprintf(w, "Error al Insertar")

	}
	var prod Pedidos
	er := json.Unmarshal(reqBody, &prod)
	if er != nil {
		fmt.Fprintf(w, "Error al Insertar el segundo")
	}
	for i := 0; i < 50000; i++ {
		if prod.Pedido[i].Fecha != "" {
			SignoMenos := regexp.MustCompile(`-`) 
			Fecha:=SignoMenos.Split(prod.Pedido[i].Fecha, -1)
			CadenaYear := Fecha[2]
			CadenaMonth := Fecha[1]
			CadenaDay := Fecha[0]
			ExA := Listas.Existe_Year(CadenaYear)
			if ExA.Year == "" {
				NoY := Listas.Node_Year{CadenaYear,Listas.NewList_Month(),nil,nil}
				Listas.Add_Year(&NoY)
			}
			var tiendass *Listas.Node_Tienda
			tiendass = Listas.Tie_Esp(prod.Pedido[i].Departamento, prod.Pedido[i].Tienda, prod.Pedido[i].Calificacion)
			if tiendass != nil {
				listProducto := Listas.NewList_Producto()
				for j := 0; j < 200; j++ {
					if prod.Pedido[i].CodigoProductos[j].Codigo != 0 {
						comprabacion := tiendass.Productos.Buscar_Producto(prod.Pedido[i].CodigoProductos[j].Codigo,0)
						if comprabacion != nil {
							Produ := Listas.Node_Producto{comprabacion,nil, nil}
							listProducto.Add_Producto(&Produ)
						}
						
					}
				}
				
				if ExA.Months == nil {
					ExA.Months = Listas.NewList_Month()
					
				}
				ExM := ExA.Months.Existe_Month(CadenaMonth)
				if ExM.Mes == "" {
					ExM1 := &Listas.Node_Month{CadenaMonth,&Listas.Matriz{nil, nil},nil,nil}
					ExA.Months.Add_Month(ExM1)
					
				}
				ExM = ExA.Months.Existe_Month(CadenaMonth)
				CadeDay, err800 := strconv.Atoi(CadenaDay)
				if err800 != nil {
					fmt.Fprintf(w, "Error al convertir el dia")
				}
				
				produc := &Listas.NodoPedido{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Dia: CadeDay, Tienda: prod.Pedido[i].Tienda, Departamento:prod.Pedido[i].Departamento,Calificacion:prod.Pedido[i].Calificacion,Productos:listProducto}
				ExM.Matriz_Dispersa.Add(produc)



			}
		}
	}
	
Listas.Calendario()

}