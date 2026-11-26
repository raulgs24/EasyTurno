package easyturno

import (
	"fmt"
	"time"
)

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

var nombresDeLosDias = map[time.Weekday]string{
	time.Monday:    "Lunes",
	time.Tuesday:   "Martes",
	time.Wednesday: "Miércoles",
	time.Thursday:  "Jueves",
	time.Friday:    "Viernes",
	time.Saturday:  "Sábado",
	time.Sunday:    "Domingo",
}

var nombresDeLasFranjas = map[Franja]string{
	FranjaDia:   "día",
	FranjaNoche: "noche",
}

func (t Turno) String() string {
	return fmt.Sprintf("%s %s", nombresDeLosDias[t.dia], nombresDeLasFranjas[t.franja])
}