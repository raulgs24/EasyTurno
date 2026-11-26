package easyturno

type Cuadrante struct {
	asignaciones map[Turno]map[*Trabajador]Puesto
}

type PuestoSinCubrir struct {
	turno  Turno
	puesto Puesto
	faltan uint
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