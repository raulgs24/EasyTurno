package domain

type Trabajador struct {
	horasContratoSemanales int
	habilidades            []Puesto
	indisponibilidad       []FranjaHoraria
}

type Puesto string

const (
	PuestoCocinero   Puesto = "Cocinero"
	PuestoRepartidor Puesto = "Repartidor"
	PuestoHorno      Puesto = "Horno"
	PuestoTelefono   Puesto = "Telefono"
)
