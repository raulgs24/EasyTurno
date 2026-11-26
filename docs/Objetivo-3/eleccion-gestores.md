# Elección de gestor de dependencias y gestor de tareas

## Gestor de dependencias

EN Go las opciones son Go Modules, [dep](https://github.com/golang/dep), [Glide](https://github.com/Masterminds/glide), [godep](https://github.com/tools/godep), [govendor](https://github.com/kardianos/govendor), [gb](https://github.com/constabulary/gb) y el modo GOPATH; dep, godep, govendor y gb están archivados, Glide avisa en su README que está prácticamente sin mantenimiento, y el modo GOPATH dejó de ser el modo por defecto en Go 1.16. Queda Go Modules, que viene con la propia orden go y se declara en go.mod

## Gestor de tareas

- Tiempo de instalación desde que empieza la instalación hasta que la orden se puede ejecutar, usando el método más rápido que ofrezca cada herramienta e instalandola en mi espacio de usuario. Es necesario medirlo porque la herramienta se instalará en cada construcción del contenedor (Objetivo 5) y en cada ejecución de la CI (Objetivo 6)
- Tiempo que añade el gestor al ejecutar check, la diferencia entre la media de <gestor> check y la de go build ./... lanzado directamente. Lo mido porque check se lanzará en cada cambio

### Opciones
Como referencia, `go build` sin gestor tarda 0.048-0.109s:
![go build sin gestor de tareas](../img/Objetivo-3/referencia-go-build.png)

Como opciones tenemos: make, Task, just, Mage y goyek
- **[make](https://www.gnu.org/software/make/)** (`Makefile`):`make check` tarda entre 0.041 y 0.109s, lo mismo que sin gestor.

  ![make check](../img/Objetivo-3/make-check.png)

- **[Task](https://github.com/go-task/task)** (`Taskfile.yml`): se tuvo que instalar; `task check` tarda entre 0.127 y 0.261s, más que sin gestor

  ![task check](../img/Objetivo-3/task-check.png)

-  **[just](https://github.com/casey/just)** (`justfile`): se tuvo que instalar, `just check` tarda entre 0.035 y 0.077s, lo mismo que sin gestor

  ![just check](../img/Objetivo-3/just-check.png)

- **[Mage](https://github.com/magefile/mage)** (`magefile.go`): se tuvo que instalar; la primera ejecución tarda 0.38s porque compila el magefile.go y las siguientes entre 0.083 y 0.087s

  ![mage check](../img/Objetivo-3/mage-check.png)

- **[goyek](https://github.com/goyek/goyek)** (biblioteca de Go): no se instala como programa, se añade con `go get`: la primera ejecución de `go run ./build check` tarda 0.39s, las siguientes entre 0.219 y 0.257s, unos 0.15-0.20s más que sin gestor.

  ![goyek check](../img/Objetivo-3/goyek-check.png)

