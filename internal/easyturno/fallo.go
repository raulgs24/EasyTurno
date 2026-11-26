package easyturno

import "fmt"

type Fallo interface {
	Descripcion() string
}

func (f PuestoSinCubrir) Descripcion() string {
	return fmt.Sprintf("%s: faltan %d en %s", f.turno, f.faltan, f.puesto)
}

func (f PuestoSinHabilidad) Descripcion() string {
	return fmt.Sprintf("%s: %s no sabe ocupar %s", f.turno, f.trabajador.nombre, f.puesto)
}

func (f TrabajadorNoDisponible) Descripcion() string {
	return fmt.Sprintf("%s: %s está en %s y no está disponible", f.turno, f.trabajador.nombre, f.puesto)
}

func (f TrabajadorSinDescanso) Descripcion() string {
	return fmt.Sprintf("%s trabaja %s y %s sin descanso", f.trabajador.nombre, f.noche, f.diaSiguiente)
}

func (f TrabajadorConExcesoDeHoras) Descripcion() string {
	return fmt.Sprintf("%s tiene %d horas asignadas y su contrato es de %d", f.trabajador.nombre, f.horasAsignadas, f.horasContrato)
}