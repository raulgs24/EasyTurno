package easyturno

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

func TestNuevoTrabajador(t *testing.T) {
	t.Run("sin horas de contrato no se puede crear", func(t *testing.T) {
		_, err := NuevoTrabajador("Pedro", 0, []Puesto{PuestoRepartidor}, nil)

		assert.IsError(t, err, ErrSinHorasContrato)
	})

	t.Run("sin ningún puesto que sepa ocupar no se puede crear", func(t *testing.T) {
		_, err := NuevoTrabajador("Pedro", 20, nil, nil)

		assert.IsError(t, err, ErrSinPuestos)
	})

	t.Run("con horas y al menos un puesto se crea sin error", func(t *testing.T) {
		_, err := NuevoTrabajador("Lucía", 20, []Puesto{PuestoHorno}, nil)

		assert.NoError(t, err)
	})

	t.Run("guarda los puestos que sabe ocupar", func(t *testing.T) {
		lucia, err := NuevoTrabajador("Lucía", 20, []Puesto{PuestoHorno, PuestoTelefono}, nil)

		assert.NoError(t, err)
		assert.Equal(t, map[Puesto]struct{}{PuestoHorno: {}, PuestoTelefono: {}}, lucia.habilidades)
	})

	t.Run("guarda los turnos en los que no está disponible", func(t *testing.T) {
		lunesNoche := Turno{dia: time.Monday, franja: FranjaNoche}

		lucia, err := NuevoTrabajador("Lucía", 20, []Puesto{PuestoHorno}, []Turno{lunesNoche})

		assert.NoError(t, err)
		assert.Equal(t, map[Turno]struct{}{lunesNoche: {}}, lucia.indisponibilidad)
	})
}