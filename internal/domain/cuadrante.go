package domain

type TurnoAsignado struct {
	puesto Puesto
	franja FranjaHoraria
}

type Cuadrante struct {
	trabajadores map[string]Trabajador
	demandas     []DemandaTurno
	asignaciones map[string][]TurnoAsignado
}

type Planificacion map[string]Cuadrante
