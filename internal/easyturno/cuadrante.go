package easyturno

type Cuadrante struct {
	asignaciones map[Turno]map[*Trabajador]Puesto
}
