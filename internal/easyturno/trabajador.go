package easyturno

type Trabajador struct {
	nombre                string
	horasMaximasSemanales uint
	habilidades           map[Puesto]struct{}
	indisponibilidad      map[Turno]struct{}
}
