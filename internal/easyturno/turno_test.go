package easyturno

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

func TestTurnosDeLaSemana(t *testing.T) {
	turnos := turnosDeLaSemana()

	t.Run("cada día tiene un turno de día y otro de noche", func(t *testing.T) {
		assert.Equal(t, 14, len(turnos))
	})

	t.Run("empieza el lunes de día y termina el domingo de noche", func(t *testing.T) {
		assert.Equal(t, Turno{dia: time.Monday, franja: FranjaDia}, turnos[0])
		assert.Equal(t, Turno{dia: time.Sunday, franja: FranjaNoche}, turnos[len(turnos)-1])
	})

	t.Run("el turno de noche va justo después del de día del mismo día", func(t *testing.T) {
		sabado := []Turno{
			{dia: time.Saturday, franja: FranjaDia},
			{dia: time.Saturday, franja: FranjaNoche},
		}
		assert.Equal(t, sabado, turnos[10:12])
	})

	t.Run("se obtiene siempre el mismo orden", func(t *testing.T) {
		assert.Equal(t, turnos, turnosDeLaSemana())
	})
}
