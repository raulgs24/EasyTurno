#Tapeo4All

##Problema
Una vez al mes quedamos todos los amigos para irnos de ruta de tapeo por el pueblo, pero hay un problema, uno de mis amigos es celiaco y otra es vegana, por lo que debemos saber cuales son los bares cuyas tapas cumplen con las restricciones alimentarias de mis amigos. Nuestro plan consiste en visitar un par de bares para tapear, como se hace de costumbre.

##De dónde viene esto
Este problema me viene de cerca puesto que uno de mis mejores amigos es alérgico al gluten y en muchos bares no hay opciones para él, por lo que solemos frecuentar siempre los mismos bares donde conocemos que hay variedades para celiacos, a pesar de que siempre nos quedamos con las ganas de probar nuevos sitios.

##Como nos organizamos
Solemos buscar a mano, establecimiento a establecimiento la carta del sitio para ver si podemos ir, pero hay veces que incluso está desactualizada y nos damos la vuelta al llegar

##Datos disponibles
Disponemos de los siguientes datos:
    + lista de miembros del grupo.
    + catálogo de bares de tapas y ubicación
    + alérgenos e ingredientes de cada bar.
    + horarios de apertura y cierre de cocina de los locales.
    + distancia máxima dispuesta a caminar entre bares
    x
##Qué hace falta procesar
Para organizar la ruta no basta buscar bares, hay que analizar la matriz de restricciones combinadas del grupo y filtrar que subconjunto de establecimientos es apto para todos a la vez
También hay que calcular distancias a pie para crear una secuencia lógica y validar que el tiempo de desplazamiento encaje con los horarios de apertura de cada local.
Si alguna cambia en tiempo real, por ejemplo, si se llena el local, se generaría una ruta alternativa desde el punto actual, que recalcule distancias y volviendo a aplicar filtros de restricciones sobre los bares restantes.

##Por qué hace falta que esté en la nube
Acutalmente la planificación cae en quien organiza el plan, esto hace que esa ruta dependa únicamente de una persona, si el grupo se divide o llegan tarde, tienen que preguntar por Whatsapp donde están. Tener la app y el itinerario accesibles en la nube permitiría que cualquiera pudiera consultar la ruta en tiempo real desde su movil y si el organizador se queda sin bateria, otro pudiera solicitar la generación de una ruta alternativa al instante

![Fotografía de la tarjeta de rol](./docs/img/tarjeta.jpeg)

*Nota: La configuración del entorno y despliegue se documentará en ficheros independientes.*
[Documentación](./docs)

