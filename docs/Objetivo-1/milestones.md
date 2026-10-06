# Milestones

## Milestone 0: Base del dominio del problema
* **Carácter:** Producto interno inicial (base para que el equipo pueda implementar la lógica de validación posterior)
* **Historias de usuario asociadas:** [HU001](historias-usuario.md).
* **Producto entregable (PMV):**
* Código fuente que plasme los elementos y las reglas del dominio del problema, tal y como se define en las historias de usuario
* **Criterio de viabilidad y validez:** 
  * El producto será viable y válido cuando el desarrollador aplique la siguiente metodología:
  * * Analizar las historias de usuario para extraes los conceptos esenciales del problema.
    * Crear los issues específicos para documentar y aislar cada tarea del desarrollo.
    * Escribir el código en ramas independientes asociadas a dichos issues.
    * Validar mediante revisión por pares en un PR, asegurando que el código refleja las reglas del negocio antes de integrarlo.

## Milestone 1: Lógica de negocio y tests
* **Carácter:** Primer producto funcional con lógica de negocio y verificable automáticamente.
* **Historias de usuario asociadas:** [HU002](historias-usuario.md) y [HU003](historias-usuario.md).
* **Producto entregable (PMV):**
* Implementación sobre el Milestone 0 que implementa la lógica de negocio necesaria para resolver los problemas planteados en las historias de usuario asociadas, junto a una batería de tests.
* **Criterio de viabilidad y validez:**
* El producto será viable y válido cuando supere tests que comprueban que el sistema da respuesta correcta a escenarios y datos definidos en dichas historias de usuario.
