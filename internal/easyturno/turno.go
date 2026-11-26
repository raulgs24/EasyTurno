package easyturno

import "time"

type Franja string

const (
	FranjaDia   Franja = "Dia"
	FranjaNoche Franja = "Noche"
)

const horasPorTurno uint = 5

type Turno struct {
	dia    time.Weekday
	franja Franja
}

var diasDeLaSemana = []time.Weekday{
	time.Monday, time.Tuesday, time.Wednesday, time.Thursday,
	time.Friday, time.Saturday, time.Sunday,
}

var franjasDelDia = []Franja{FranjaDia, FranjaNoche}

func turnosDeLaSemana() []Turno {
	turnos := make([]Turno, 0, len(diasDeLaSemana)*len(franjasDelDia))
	for _, dia := range diasDeLaSemana {
		for _, franja := range franjasDelDia {
			turnos = append(turnos, Turno{dia: dia, franja: franja})
		}
	}
	return turnos
}
