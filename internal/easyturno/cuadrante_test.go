package easyturno

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

func TestPuestosSinCubrir(t *testing.T) {
	lucia := &Trabajador{nombre: "Lucía"}
	pedro := &Trabajador{nombre: "Pedro"}
	lunesDia := Turno{dia: time.Monday, franja: FranjaDia}
	lunesNoche := Turno{dia: time.Monday, franja: FranjaNoche}
	sabadoNoche := Turno{dia: time.Saturday, franja: FranjaNoche}
	domingoDia := Turno{dia: time.Sunday, franja: FranjaDia}

	t.Run("con las personas que pide la demanda no falta nadie", func(t *testing.T) {
		demanda := Demanda{plazas: map[Turno]map[Puesto]uint{
			sabadoNoche: {PuestoHorno: 1, PuestoTelefono: 1},
		}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			sabadoNoche: {lucia: PuestoHorno, pedro: PuestoTelefono},
		}}

		assert.Zero(t, cuadrante.PuestosSinCubrir(demanda))
	})

	t.Run("el sábado de noche sin nadie en el horno", func(t *testing.T) {
		demanda := Demanda{plazas: map[Turno]map[Puesto]uint{
			sabadoNoche: {PuestoHorno: 1, PuestoTelefono: 1},
		}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			sabadoNoche: {pedro: PuestoTelefono},
		}}

		assert.Equal(t,
			[]PuestoSinCubrir{{turno: sabadoNoche, puesto: PuestoHorno, faltan: 1}},
			cuadrante.PuestosSinCubrir(demanda))
	})

	t.Run("se necesitan dos personas al teléfono y solo hay una", func(t *testing.T) {
		demanda := Demanda{plazas: map[Turno]map[Puesto]uint{
			sabadoNoche: {PuestoTelefono: 2},
		}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			sabadoNoche: {pedro: PuestoTelefono},
		}}

		assert.Equal(t,
			[]PuestoSinCubrir{{turno: sabadoNoche, puesto: PuestoTelefono, faltan: 1}},
			cuadrante.PuestosSinCubrir(demanda))
	})

	t.Run("más personas de las necesarias no es un puesto sin cubrir", func(t *testing.T) {
		demanda := Demanda{plazas: map[Turno]map[Puesto]uint{
			sabadoNoche: {PuestoRepartidor: 1},
		}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			sabadoNoche: {lucia: PuestoRepartidor, pedro: PuestoRepartidor},
		}}

		assert.Zero(t, cuadrante.PuestosSinCubrir(demanda))
	})

	t.Run("los huecos salen en el orden de los turnos y de los puestos", func(t *testing.T) {
		demanda := Demanda{plazas: map[Turno]map[Puesto]uint{
			domingoDia: {PuestoRepartidor: 1},
			lunesNoche: {PuestoTelefono: 1},
			lunesDia:   {PuestoHorno: 1, PuestoCocinero: 2},
		}}
		cuadrante := Cuadrante{}

		assert.Equal(t,
			[]PuestoSinCubrir{
				{turno: lunesDia, puesto: PuestoCocinero, faltan: 2},
				{turno: lunesDia, puesto: PuestoHorno, faltan: 1},
				{turno: lunesNoche, puesto: PuestoTelefono, faltan: 1},
				{turno: domingoDia, puesto: PuestoRepartidor, faltan: 1},
			},
			cuadrante.PuestosSinCubrir(demanda))
	})
}
