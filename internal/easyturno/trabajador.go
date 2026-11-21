package easyturno

import "time"

type Trabajador struct {
	nombre                 string
	horasContratoSemanales int
	habilidades            map[Puesto]struct{}
	indisponibilidad       map[time.Weekday]struct{}
}
