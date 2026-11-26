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

func TestPuestosSinHabilidad(t *testing.T) {
	lucia := &Trabajador{nombre: "Lucía", habilidades: map[Puesto]struct{}{PuestoHorno: {}, PuestoTelefono: {}}}
	pedro := &Trabajador{nombre: "Pedro", habilidades: map[Puesto]struct{}{PuestoRepartidor: {}}}
	viernesNoche := Turno{dia: time.Friday, franja: FranjaNoche}
	sabadoDia := Turno{dia: time.Saturday, franja: FranjaDia}

	t.Run("si cada uno está en un puesto que sabe ocupar no hay fallo", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesNoche: {lucia: PuestoHorno, pedro: PuestoRepartidor},
		}}

		assert.Zero(t, cuadrante.PuestosSinHabilidad())
	})

	t.Run("Pedro, que solo sabe repartir, en el horno el viernes por la noche", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesNoche: {pedro: PuestoHorno},
		}}

		assert.Equal(t,
			[]PuestoSinHabilidad{{turno: viernesNoche, trabajador: pedro, puesto: PuestoHorno}},
			cuadrante.PuestosSinHabilidad())
	})

	t.Run("quien sabe ocupar varios puestos puede estar en cualquiera", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesNoche: {lucia: PuestoTelefono},
			sabadoDia:    {lucia: PuestoHorno},
		}}

		assert.Zero(t, cuadrante.PuestosSinHabilidad())
	})

	t.Run("los fallos salen en el orden de los turnos y, en cada turno, por nombre", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			sabadoDia:    {pedro: PuestoTelefono, lucia: PuestoRepartidor},
			viernesNoche: {pedro: PuestoCocinero},
		}}

		assert.Equal(t,
			[]PuestoSinHabilidad{
				{turno: viernesNoche, trabajador: pedro, puesto: PuestoCocinero},
				{turno: sabadoDia, trabajador: lucia, puesto: PuestoRepartidor},
				{turno: sabadoDia, trabajador: pedro, puesto: PuestoTelefono},
			},
			cuadrante.PuestosSinHabilidad())
	})
}

func TestTrabajadoresNoDisponibles(t *testing.T) {
	lunesDia := Turno{dia: time.Monday, franja: FranjaDia}
	lunesNoche := Turno{dia: time.Monday, franja: FranjaNoche}
	martesNoche := Turno{dia: time.Tuesday, franja: FranjaNoche}
	pedro := &Trabajador{nombre: "Pedro"}

	t.Run("si todos trabajan cuando están disponibles no hay fallo", func(t *testing.T) {
		lucia := &Trabajador{nombre: "Lucía", indisponibilidad: map[Turno]struct{}{lunesNoche: {}}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			martesNoche: {lucia: PuestoTelefono, pedro: PuestoHorno},
		}}

		assert.Zero(t, cuadrante.TrabajadoresNoDisponibles())
	})

	t.Run("Lucía no puede el lunes por la noche y está en el teléfono", func(t *testing.T) {
		lucia := &Trabajador{nombre: "Lucía", indisponibilidad: map[Turno]struct{}{lunesNoche: {}}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			lunesNoche: {lucia: PuestoTelefono},
		}}

		assert.Equal(t,
			[]TrabajadorNoDisponible{{turno: lunesNoche, trabajador: lucia, puesto: PuestoTelefono}},
			cuadrante.TrabajadoresNoDisponibles())
	})

	t.Run("no poder el lunes por la noche no impide trabajar el lunes de día", func(t *testing.T) {
		lucia := &Trabajador{nombre: "Lucía", indisponibilidad: map[Turno]struct{}{lunesNoche: {}}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			lunesDia: {lucia: PuestoTelefono},
		}}

		assert.Zero(t, cuadrante.TrabajadoresNoDisponibles())
	})

	t.Run("si no puede en todo el lunes falla de día y de noche", func(t *testing.T) {
		lucia := &Trabajador{nombre: "Lucía", indisponibilidad: map[Turno]struct{}{lunesDia: {}, lunesNoche: {}}}
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			lunesNoche: {lucia: PuestoHorno},
			lunesDia:   {lucia: PuestoTelefono, pedro: PuestoRepartidor},
		}}

		assert.Equal(t,
			[]TrabajadorNoDisponible{
				{turno: lunesDia, trabajador: lucia, puesto: PuestoTelefono},
				{turno: lunesNoche, trabajador: lucia, puesto: PuestoHorno},
			},
			cuadrante.TrabajadoresNoDisponibles())
	})
}


func TestTrabajadoresSinDescanso(t *testing.T) {
	lucia := &Trabajador{nombre: "Lucía"}
	pedro := &Trabajador{nombre: "Pedro"}
	lunesDia := Turno{dia: time.Monday, franja: FranjaDia}
	viernesDia := Turno{dia: time.Friday, franja: FranjaDia}
	viernesNoche := Turno{dia: time.Friday, franja: FranjaNoche}
	sabadoDia := Turno{dia: time.Saturday, franja: FranjaDia}
	sabadoNoche := Turno{dia: time.Saturday, franja: FranjaNoche}
	domingoNoche := Turno{dia: time.Sunday, franja: FranjaNoche}

	t.Run("de noche el viernes y de día el sábado se queda sin descanso", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesNoche: {lucia: PuestoHorno},
			sabadoDia:    {lucia: PuestoTelefono},
		}}

		assert.Equal(t,
			[]TrabajadorSinDescanso{{trabajador: lucia, noche: viernesNoche, diaSiguiente: sabadoDia}},
			cuadrante.TrabajadoresSinDescanso())
	})

	t.Run("el turno de día y el de noche del mismo día no es falta de descanso", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesDia:   {lucia: PuestoHorno},
			viernesNoche: {lucia: PuestoTelefono},
		}}

		assert.Zero(t, cuadrante.TrabajadoresSinDescanso())
	})

	t.Run("de noche dos días seguidos no es falta de descanso", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesNoche: {lucia: PuestoHorno},
			sabadoNoche:  {lucia: PuestoHorno},
		}}

		assert.Zero(t, cuadrante.TrabajadoresSinDescanso())
	})

	t.Run("si el día siguiente trabaja otra persona no hay fallo", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			viernesNoche: {lucia: PuestoHorno},
			sabadoDia:    {pedro: PuestoHorno},
		}}

		assert.Zero(t, cuadrante.TrabajadoresSinDescanso())
	})

	t.Run("la noche del domingo no se compara con el lunes de la misma semana", func(t *testing.T) {
		cuadrante := Cuadrante{asignaciones: map[Turno]map[*Trabajador]Puesto{
			domingoNoche: {lucia: PuestoHorno},
			lunesDia:     {lucia: PuestoHorno},
		}}

		assert.Zero(t, cuadrante.TrabajadoresSinDescanso())
	})
}