package easyturno

import "time"

type semana struct {
	anio   int
	numero int
}

type CuadranteSemanal struct {
	semana       semana
	asignaciones map[time.Weekday][]AsignacionTurno
}
