package easyturno

import "errors"

var (
	ErrSinHorasContrato = errors.New("el trabajador no tiene horas de contrato")
	ErrSinPuestos       = errors.New("el trabajador no sabe ocupar ningún puesto")
)

type Trabajador struct {
	nombre                string
	horasMaximasSemanales uint
	habilidades           map[Puesto]struct{}
	indisponibilidad      map[Turno]struct{}
}

func NewTrabajador(nombre string, horasMaximasSemanales uint, habilidades []Puesto, indisponibilidad []Turno) (*Trabajador, error) {
	if horasMaximasSemanales == 0 {
		return nil, ErrSinHorasContrato
	}
	if len(habilidades) == 0 {
		return nil, ErrSinPuestos
	}

	trabajador := &Trabajador{
		nombre:                nombre,
		horasMaximasSemanales: horasMaximasSemanales,
		habilidades:           make(map[Puesto]struct{}, len(habilidades)),
		indisponibilidad:      make(map[Turno]struct{}, len(indisponibilidad)),
	}
	for _, puesto := range habilidades {
		trabajador.habilidades[puesto] = struct{}{}
	}
	for _, turno := range indisponibilidad {
		trabajador.indisponibilidad[turno] = struct{}{}
	}
	return trabajador, nil
}