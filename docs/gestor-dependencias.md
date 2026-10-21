# Elección del gestor de dependencias

Raúl Gabaldón Sánchez

Para la elección del gestor de dependencias, nos tenemos que hacer las siguientes preguntas:

## ¿Go define cómo se declaran las dependencias y la versión del lenguaje? ¿Dónde lo define?

Sí, Go lo define en el propio go.mod. Cada módulo de go se define con un go.mod que declara:

- module: ruta del módulo
- go: la versión mínima de Go necesaria para compilar
- require: dependencias y su versión minima (en el go.mod no lo ponemos porque empleamos biblioteca estándar)

Fuentes: [go.dev/doc/modules/gomod-ref](https://go.dev/doc/modules/gomod-ref), [go.dev/ref/mod](https://go.dev/ref/mod)

## ¿La herramienta sigue mantenida o está archivada?

Tenemos como opciones godep, glide, gb, govendor, dep, Go Modules y GOPATH + go get.

De estos, godep, gb, govendor y dep están archivados. glide no está archivado formalmente pero se declara casi sin mantenimiento. GOPATH + go get todavía lo conserva por compatibilidad, pero no es el modo por defecto ni el recomendado.

Ya solo nos queda una opción, **Go Modules**, es más, dentro de dep, glide y govendor se recomienda migrar a Go Modules, el cual es mantenido en la actualidad por el propio Go

Fuentes: [github.com/golang/dep](https://github.com/golang/dep), [github.com/Masterminds/glide](https://github.com/Masterminds/glide), [github.com/tools/godep](https://github.com/tools/godep), [github.com/kardianos/govendor](https://github.com/kardianos/govendor), [github.com/constabulary/gb](https://github.com/constabulary/gb), [go.dev/doc/go1.16](https://go.dev/doc/go1.16)

## ¿Con cero dependencias externas, que necesito realmente?

Lo único que necesita mi proyecto para compilar es Go, en la versión que marca el go.mod, nada más, puesto que: el código solo importa time, que es de la biblioteca estándar y viene incluida en la instalación de Go. El go.mod no tiene el bloque require y no existe go.sum porque no hay nada que verificar (Este fichero solo aparece cuando se añade una dependencia).

Mirando a futuro, si es posible que existan ciertas dependencias y se deba crear esos archivos, pero no se pueden inventar dependencias que se tienen, el criterio se aplica al proyecto tal y como está

## ¿Verifica la integridad de lo que descarga?

Sí, verifica la integridad de todo lo que descarga a dos niveles:

En go.sum (verificación local) donde para cada dependencia guarda un hash criptográfico. Cada vez que la orden go descarga o usa ese módulo, recalcula el hash y lo compara con el guardado. Si no coincide falla y no usa ese código. Así se puede ver si alguien altera una versión ya publicada.

Por otra parte, se verifica en la base de datos de checksums, cuando una dependencia aparece por primera vez y aún no está en el go.sum. Go consulta esta base de datos pública, donde cada versión tiene un único hash que ve todo el mundo. Así no te fias solo de lo que entrega el servidor de descarga.

Fuentes: [go.dev/ref/mod](https://go.dev/ref/mod)

## ¿Hay que instalar algo además de Go?

Para Go Modules no, viene integrado con la orden go y está activado por defecto. Las otras alternativas en cambio eran programas externos instalados por separado.

## Conclusión

Go Modules es la única opción que es estándar, está mantenida y no añade ninguna instalación. Y con cero dependencias externas, cualquier herramienta adicional sería coste sin beneficio

