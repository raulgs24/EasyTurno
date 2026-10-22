package domain

type DemandaTurno struct {
	franja            FranjaHoraria
	puestosRequeridos map[Puesto]int
}
