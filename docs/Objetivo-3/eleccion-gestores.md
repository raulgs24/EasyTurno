# Elección de gestor de dependencias y gestor de tareas

## Gestor de dependencias

EasyTurno solo importa `time`, de la biblioteca estándar, así que no tiene dependencias externas. Aun así, necesito:

- **Que `internal/easyturno` compile**, porque `check` depende de ello.
- **Que compile igual en mi máquina, en el contenedor (objetivo 5) y en la CI (objetivo 6).**
- **No añadir peso ni tiempo de instalación a la imagen.**
- **Que cuando llegue la primera dependencia (por ejemplo, en los tests) se instale siempre la misma versión.**

Las alternativas a tener en cuenta son Go Modules, dep, Glide y el modo GOPATH.

Descartadas:

- **dep y Glide:** archivados.
- **Modo GOPATH:** sigue en la herramienta, pero no fija versiones.

Por tanto, la decisión es Go Modules: es el estándar del lenguaje, el modo por defecto desde Go 1.16, y cubre las cuatro necesidades anteriores:

- **Compila:** `go.mod` declara el módulo; sin él, `go build` falla.
- **Igual en todas partes:** la línea `go` de `go.mod` fija la versión de Go.
- **Sin peso:** va dentro de la orden `go`, no instala nada.
- **Misma versión:** `go.mod` fija versiones y `go.sum` guarda sus sumas de verificación.

## Gestor de tareas

Go Modules no sirve como gestor de tareas, ya que la orden go no tiene un fichero donde definir tareas propias. Necesito:
- **Una orden única para cada tarea:** ahora `install` y `check`
- **Que check compruebe solo la sintaxis**, es decir, que compile
- **Que se instale rápido,** ya que se instalará en cada construcción del contenedor (objetivo 5) y cada ejecución de la CI (objetivo 6)
- **Que no añada tiempo apreciable al ejecutar las tareas,** porque se lanzarán en cada cambio
- **Que no añada dependencias a go.mod.**

### Criterios:
* Tiempo de instalación.
* Tiempo de ejecución de `check` comparado con `go build`.
* Dependencias que añade a `go.mod`.

### Opciones
Como referencia, `go build` sin gestor tarda 0.048-0.109s:
![go build sin gestor de tareas](../img/Objetivo-3/referencia-go-build.png)

Como opciones tenemos: make, Task, just, Mage y goyek
- **[make]** (`Makefile`): Ya está instalado en mi máquina; `make check` tarda entre 0.041 y 0.109s, lo mismo que sin gestor.

  ![make ya instalado](../img/Objetivo-3/make-instalacion.png)
  ![make check](../img/Objetivo-3/make-check.png)

- **[Task]** (`Taskfile.yml`): instalarlo con `go install` que lo compila, tarda 1 min 4s; `task check` tarda entre 0.127 y 0.261s, unos 0.05s-0.08s mas que sin gestor.

  ![Instalación de Task](../img/Objetivo-3/task-instalacion.png)
  ![task check](../img/Objetivo-3/task-check.png)

- **[just]** (`justfile`): instalarlo con `apt` tarda 8,305s; `just check` tarda entre 0.035 y 0.077s, lo mismo que sin gestor.

  ![Instalación de just](../img/Objetivo-3/just-instalacion.png)
  ![just check](../img/Objetivo-3/just-check.png)

- **[Mage]** (`magefile.go`): instalarlo con go install tarda 9.15s; la primera ejecución tarda 0.38s porque compila el magefile.go y las siguientes entre 0.083 y 0.087s

  ![Instalación de Mage](../img/Objetivo-3/mage-instalacion.png)
  ![mage check](../img/Objetivo-3/mage-check.png)

- **[goyek]** (biblioteca de Go): añadirlo con `go get` tarde 1.95s: la primera ejecución de `go run ./build check` tarda 0.39s, las siguientes entre 0.219 y 0.257s, unos 0.15-0.20s más que sin gestor; añade una dependecia a `go.mod` y crea `go.sum`

  ![Instalación de goyek](../img/Objetivo-3/goyek-instalacion.png)
  ![goyek check](../img/Objetivo-3/goyek-check.png)
  ![go.mod con goyek](../img/Objetivo-3/goyek-gomod.png)


Por tanto, los criterios apuntan a **make**: es el único que no hubo que instalar, no añade tiempo a check y no toca go.mod. Just empata en tiempo de ejecucuón, y en go.mod, simplemente pierde en la instalación.

