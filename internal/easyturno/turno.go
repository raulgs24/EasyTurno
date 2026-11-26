package easyturno

import "time"

type Franja string

const (
	FranjaDia   Franja = "Dia"
	FranjaNoche Franja = "Noche"
)

type Turno struct {
	dia    time.Weekday
	franja Franja
}
