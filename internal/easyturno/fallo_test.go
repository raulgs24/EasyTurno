package easyturno

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

func TestDescripcion(t *testing.T) {
	lucia := &Trabajador{nombre: "Lucía"}
	viernesNoche := Turno{dia: time.Friday, franja: FranjaNoche}
	sabadoDia := Turno{dia: time.Saturday, franja: FranjaDia}

	t.Run("puesto sin cubrir", func(t *testing.T) {
		fallo := PuestoSinCubrir{turno: viernesNoche, puesto: PuestoHorno, faltan: 1}

		assert.Equal(t, "Viernes noche: faltan 1 en Horno", fallo.Descripcion())
	})

	t.Run("puesto sin habilidad", func(t *testing.T) {
		fallo := PuestoSinHabilidad{turno: viernesNoche, trabajador: lucia, puesto: PuestoRepartidor}

		assert.Equal(t, "Viernes noche: Lucía no sabe ocupar Repartidor", fallo.Descripcion())
	})

	t.Run("trabajador no disponible", func(t *testing.T) {
		fallo := TrabajadorNoDisponible{turno: sabadoDia, trabajador: lucia, puesto: PuestoTelefono}

		assert.Equal(t, "Sábado día: Lucía está en Telefono y no está disponible", fallo.Descripcion())
	})

	t.Run("trabajador sin descanso", func(t *testing.T) {
		fallo := TrabajadorSinDescanso{trabajador: lucia, noche: viernesNoche, diaSiguiente: sabadoDia}

		assert.Equal(t, "Lucía trabaja Viernes noche y Sábado día sin descanso", fallo.Descripcion())
	})

	t.Run("trabajador con exceso de horas", func(t *testing.T) {
		fallo := TrabajadorConExcesoDeHoras{trabajador: lucia, horasAsignadas: 15, horasContrato: 10}

		assert.Equal(t, "Lucía tiene 15 horas asignadas y su contrato es de 10", fallo.Descripcion())
	})
}