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
Como referencia, `go build` sin gestor tarda:
![go build sin gestor de tareas](../img/Objetivo-3/referencia-go-build.png)

Como opciones tenemos: make, Task, just, Mage y goyek
- **[make]** (`Makefile`): Ya está instalado en mi máquina; `make check` tarda entre 0.04 y 0.08s, lo mismo que sin gestor.

  ![make ya instalado](../img/Objetivo-3/make-instalacion.png)
  ![make check](../img/Objetivo-3/make-check.png)

- **[Task]** (`Taskfile.yml`): instalarlo con `go install` que lo compila, tarda 1 min 4s; `task check` tarda entre 0.1 y 0.16s, unos 0.05s-0.08s mas que sin gestor.

  ![Instalación de Task](../img/Objetivo-3/task-instalacion.png)
  ![task check](../img/Objetivo-3/task-check.png)

- **[just]** (`justfile`): instalarlo con `apt` tarda 8,305s; `just check` tarda entre 0.04 y 0.07s, lo mismo que sin gestor.

  ![Instalación de just](../img/Objetivo-3/just-instalacion.png)
  ![just check](../img/Objetivo-3/just-check.png)
