#Milestones

Cada milestone define un Producto Mínimamente Viable (PMV) entregable, que evoluciona de forma gradual hacia la solución de las [historias de usuario](historias-usuario.md)

##Milestone 0: Modelo de los datos
* **Carácter:** Producto interno inicial (base para que el equipo pueda implementar la lógica de validación posterior)
* **Historias de usuario asociadas:** [HU001](historias-usuario.md).
* **Producto entregable (PMV):**
* Módulo que representa las estructuras de datos inmutables de la pizzería:
* * Los trabajadores de la plantilla con sus puestos capacitados (cocina, horno, teléfono y reparto), sus horas máximas de contrato y franjas de indisponibilidad.
  * La configuración de los turnos diarios y demanda mínima de puestos según el día de la semana.

##Milestone 1: Validación de restricciones y reajuste de cuadrantes
* **Carácter:** Primer producto funcional con lógica de negocio.
* **Historias de usuario asociadas:** [HU002](historias-usuario.md) y [HU003](historias-usuario.md).
* **Producto entregable (PMV):**
* Empleo de las estructuras del Milestone 0 para comprobar si un cuadrante cumple todas las reglas operativas y laborales:
* * Validación de cobertura y restricciones.
  * Reajuste de urgencia ante bajas imprevistas.
