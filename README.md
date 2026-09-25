# EasyTurno

## Problema
La pizzeria donde trabajo tiene un problema a la hora de cuadrar los horarios, puesto que supone un rompecabezas para mi jefe, no basta con asignar huecos,un dia normal se necesitan 4 cocineros, 1 en el horno, 1 atendiendo a los teléfonos y 4 repartidores, respetando las horas de descanso y las horas por contrato, pero los findes la demanda sube haciendo que se necesiten 7 cocineros,1 de ellos en el horno, 2 atendiendo telefonos y 5 repartidores. También si alguien se da de baja hay que reestructurar el cuadrante.

## Como se organiza
Se suele realizar el cuadrante mediante ensayo y error, verificando de forma visual que se cumplen las condiciones necesarias, lo que puede dar fallos inesperados. 

## Datos disponibles
El sistema dispondrá de un csv con los datos.
Disponemos de los siguientes datos:
    + Plantilla de empleados
    + Disponibilidad de empleados
    + Curva de demanda esperada
    
## Qué hace falta procesar
Hay que resolver un problema con una serie de restricciones, las cuales son
    + Calcular permutaciones de asignación de turnos y descartar aquellas que violen reglas estrictas
    + Asegurar que los cuadrantes resultantes cumplen estrictamente con las habilidades requeridas en cada dia.
    + Ante una incidencia en tiempo real, el sistema debe destruir la asignación actual y computar en segundos una nueva distribución de emergencia
    

![Fotografía de la tarjeta de rol](./docs/img/tarjeta.jpeg)
                                                                       
*Nota: La configuración del entorno y despliegue se documentará en ficheros independientes.*
[Documentación](./docs)

