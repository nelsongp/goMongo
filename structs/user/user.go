package user

import "time"

//estructura de objeto usuario
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//funcion de dar alta a usuario
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
