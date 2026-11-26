package easyturno

type Puesto string

const (
	PuestoCocinero   Puesto = "Cocinero"
	PuestoRepartidor Puesto = "Repartidor"
	PuestoHorno      Puesto = "Horno"
	PuestoTelefono   Puesto = "Telefono"
)

var puestosDelLocal = []Puesto{PuestoCocinero, PuestoRepartidor, PuestoHorno, PuestoTelefono}
