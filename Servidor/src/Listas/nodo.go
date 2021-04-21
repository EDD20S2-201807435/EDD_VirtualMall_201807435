package Listas

type NodoB struct{
	Max int
	NodoPadre *NodoB
	Keys []*Key
}
func NewNodoB(max int)*NodoB{
	keys :=make([]*Key,max)
	n:=NodoB{max,nil,keys}
	return &n
}
func (this *NodoB)Colocar(i int, llave *Key){
	this.Keys[i] = llave
} 
