package easyturno

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

type Cuadrante struct {
	asignaciones map[Turno]map[*Trabajador]Puesto
}

type PuestoSinCubrir struct {
	turno  Turno
	puesto Puesto
	faltan uint
}

type PuestoSinHabilidad struct {
	turno      Turno
	trabajador *Trabajador
	puesto     Puesto
}

type TrabajadorConExcesoDeHoras struct {
	trabajador     *Trabajador
	horasAsignadas uint
	horasContrato  uint
}


func (c Cuadrante) Revisar(demanda Demanda) []Fallo {
	return slices.Concat(
		comoFallos(c.PuestosSinCubrir(demanda)),
		comoFallos(c.PuestosSinHabilidad()),
		comoFallos(c.TrabajadoresNoDisponibles()),
		comoFallos(c.TrabajadoresSinDescanso()),
		comoFallos(c.TrabajadoresConExcesoDeHoras()),
	)
}

func comoFallos[T Fallo](lista []T) []Fallo {
	fallos := make([]Fallo, 0, len(lista))
	for _, fallo := range lista {
		fallos = append(fallos, fallo)
	}
	return fallos
}

func (c Cuadrante) PuestosSinCubrir(demanda Demanda) []PuestoSinCubrir {
	var sinCubrir []PuestoSinCubrir
	for _, turno := range turnosDeLaSemana() {
		asignadas := c.personasPorPuesto(turno)
		for _, puesto := range puestosDelLocal {
			necesarias := demanda.plazas[turno][puesto]
			if asignadas[puesto] < necesarias {
				sinCubrir = append(sinCubrir, PuestoSinCubrir{
					turno:  turno,
					puesto: puesto,
					faltan: necesarias - asignadas[puesto],
				})
			}
		}
	}
	return sinCubrir
}

func (c Cuadrante) personasPorPuesto(turno Turno) map[Puesto]uint {
	personas := make(map[Puesto]uint)
	for _, puesto := range c.asignaciones[turno] {
		personas[puesto]++
	}
	return personas
}

func (c Cuadrante) PuestosSinHabilidad() []PuestoSinHabilidad {
	var sinHabilidad []PuestoSinHabilidad
	for _, turno := range turnosDeLaSemana() {
		for _, trabajador := range c.trabajadoresDelTurno(turno) {
			puesto := c.asignaciones[turno][trabajador]
			if !trabajador.sabeOcupar(puesto) {
				sinHabilidad = append(sinHabilidad, PuestoSinHabilidad{
					turno:      turno,
					trabajador: trabajador,
					puesto:     puesto,
				})
			}
		}
	}
	return sinHabilidad
}

func (c Cuadrante) trabajadoresDelTurno(turno Turno) []*Trabajador {
	return ordenadosPorNombre(maps.Keys(c.asignaciones[turno]))
}

func ordenadosPorNombre(trabajadores iter.Seq[*Trabajador]) []*Trabajador {
	return slices.SortedFunc(trabajadores, func(a, b *Trabajador) int {
		return cmp.Compare(a.nombre, b.nombre)
	})
}

type TrabajadorNoDisponible struct {
	turno      Turno
	trabajador *Trabajador
	puesto     Puesto
}

func (c Cuadrante) TrabajadoresNoDisponibles() []TrabajadorNoDisponible {
	var noDisponibles []TrabajadorNoDisponible
	for _, turno := range turnosDeLaSemana() {
		for _, trabajador := range c.trabajadoresDelTurno(turno) {
			if !trabajador.disponible(turno) {
				noDisponibles = append(noDisponibles, TrabajadorNoDisponible{
					turno:      turno,
					trabajador: trabajador,
					puesto:     c.asignaciones[turno][trabajador],
				})
			}
		}
	}
	return noDisponibles
}


type TrabajadorSinDescanso struct {
	trabajador   *Trabajador
	noche        Turno
	diaSiguiente Turno
}


func (c Cuadrante) TrabajadoresSinDescanso() []TrabajadorSinDescanso {
	var sinDescanso []TrabajadorSinDescanso
	for i := range len(diasDeLaSemana) - 1 {
		noche := Turno{dia: diasDeLaSemana[i], franja: FranjaNoche}
		diaSiguiente := Turno{dia: diasDeLaSemana[i+1], franja: FranjaDia}
		for _, trabajador := range c.trabajadoresDelTurno(noche) {
			if _, trabaja := c.asignaciones[diaSiguiente][trabajador]; trabaja {
				sinDescanso = append(sinDescanso, TrabajadorSinDescanso{
					trabajador:   trabajador,
					noche:        noche,
					diaSiguiente: diaSiguiente,
				})
			}
		}
	}
	return sinDescanso
}


func (c Cuadrante) TrabajadoresConExcesoDeHoras() []TrabajadorConExcesoDeHoras {
	var conExceso []TrabajadorConExcesoDeHoras
	turnos := c.turnosPorTrabajador()
	for _, trabajador := range ordenadosPorNombre(maps.Keys(turnos)) {
		horas := turnos[trabajador] * horasPorTurno
		if horas > trabajador.horasMaximasSemanales {
			conExceso = append(conExceso, TrabajadorConExcesoDeHoras{
				trabajador:     trabajador,
				horasAsignadas: horas,
				horasContrato:  trabajador.horasMaximasSemanales,
			})
		}
	}
	return conExceso
}

func (c Cuadrante) turnosPorTrabajador() map[*Trabajador]uint {
	turnos := make(map[*Trabajador]uint)
	for _, asignaciones := range c.asignaciones {
		for trabajador := range asignaciones {
			turnos[trabajador]++
		}
	}
	return turnos
}