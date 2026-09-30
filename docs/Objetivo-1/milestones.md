# Milestones

Cada milestone define un Producto Mínimamente Viable (PMV) entregable, que evoluciona de forma gradual hacia la solución de las [historias de usuario](historias-usuario.md)

## Milestone 0: Base del dominio del problema
* **Carácter:** Producto interno inicial (base para que el equipo pueda implementar la lógica de validación posterior)
* **Historias de usuario asociadas:** [HU001](historias-usuario.md).
* **Producto entregable (PMV):**
* Módulo que representa las estructuras de datos inmutables de la pizzería:
* * Los trabajadores de la plantilla con sus puestos capacitados (cocina, horno, teléfono y reparto), sus horas máximas de contrato y franjas de indisponibilidad.
  * La configuración de los turnos diarios y demanda mínima de puestos según el día de la semana.
* **Criterio de viabilidad y validez:**
* El producto será viable y válido cuando cada cambio responda a un issue específico derivado del análisis de las historias de usuario, verificando que refleja fielmente el problema planteado

## Milestone 1: Lógica de negocio y tests
* **Carácter:** Primer producto funcional con lógica de negocio y verificable automáticamente.
* **Historias de usuario asociadas:** [HU002](historias-usuario.md) y [HU003](historias-usuario.md).
* **Producto entregable (PMV):**
* Implementación sobre el Milestone 0 que implementa la lógica de negocio necesaria para resolver los problemas planteados en las historias de usuario asociadas, junto a una batería de tests.
* **Criterio de viabilidad y validez:**
* El producto será viable y válido cuando supere tests que comprueban que el sistema da respuesta correcta a escenarios y datos definidos en dichas historias de usuario.
