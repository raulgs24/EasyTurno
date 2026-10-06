# Milestones
Todas las milestones avanzan en resolución del problema de la [HU001](historias-usuario.md)
La metodología que sigue es el Desarrollo Guiado por Comportamiento (de la HU001 se extraen escenarios, de cada escenario surgen issues y cada commit referencia el issue al que responde)

## Milestone 0
* **Producto:** Un paquete instalable del lenguaje elegido, que el M1 podrá importar y usar directamente.
* **Validez:** El código passa sin errores el compilador o comprobador de sintaxis del lenguaje (a elegir en un futuro).Se comprueba mientras se desarrolla, no al final. Cada cambio llega en un pr que responde a un issue surgido de un escenario de la HU001, y se revisa antes del merge comprobando que el codigo responde al issue

## Milestone 1
* **Producto:** Una nueva versión del paquete de la m0, que incorpora el código capaz de dar respuesta al problema de la HU001, junto a una batería de test.
* **Validez:** Los tests, escritos a partir de los escenarios de la HU001, se ejecutan automáticamente y pasan. Se mantiene el mismo proceso de issue -> commit -> revisión en el PR
