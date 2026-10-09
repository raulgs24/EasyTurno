# Milestones
Todas las milestones avanzan en la resolución del problema de la [HU001](historias-usuario.md)

## Milestone 0
Se empleará DDD para trabajar sobre la HU001, a partir del cual se comprenderá el problema planteado de los turnos de la pizzería. Esto será válido cuando el código use los mismos términos y reglas que Carlos emplea al hacer el cuadrante (trabajadores, puestos, turnos, descansos y contratos), de forma que el problema de la HU001 se reconozca al leer codigo.

## Milestone 1
Carlos pierde mucho tiempo haciendo cuadrantes sin saber si cubre todos los puestos, si respeta los horarios de cada persona y si cumple los horarios legales de descanso. Partiendo de la m0, el problema se divide en preguntas más pequeñas (comprobables mediante tests), por ejemplo, ¿queda algún dia sin nadie en el horno? o ¿alguien trabaja dos turnos seguidos sin su descanso? ,cada pregunta va en un issue, cada issue se resuelve con código sobre m0. Este milestone será válido siempre y cuando ningún issue se cierre sin que su test pase, esos test escritos a través de escenarios de la HU001.
