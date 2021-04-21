package Listas

type Key struct {
	Dpi      int64 
	Nombre      string
	Correo string
	Password      string
	Cuenta    string
	Izquierdo *NodoB
	Derecho *NodoB
}
func NewKey(Dpi int64, Nombre string,Correo string,Password string,Cuenta string)*Key{
	k:= Key{Dpi,Nombre,Correo,Password,Cuenta,nil,nil}
	return &k
}