# Elección de herramientas de test

Ningún issue de la HU001 se cierra sin un test que lo compruebe. Pasa eso necesito:

* **Escribir comprobaciones:** que comparen lo que devuelve el código con lo esperado; muchas veces, listas de fallos del cuadrante
* **Agrupar los casos de cada issue:** preparar sus datos (plantilla y cuadrante) en cada test
* **Ejecutarlos con una sola orden:** igual en mi máquina, en el contenedor y en la CI

## Niveles
Son tres niveles distintos y se eligen por separado:

- **Biblioteca de aserciones:** compara el valor obtenido con el esperado y, si no coincide, da un mensaje y marca el test como fallido.
- **Framework de test:** organiza los tests y subtests, y prepara o limpia datos antes y después de cada uno.
- **Herramienta CLI:** encuentra los tests, los ejecuta y da el resultado

## Criterios
Los criterios fijados antes de mirar las opciones son:
1. **Dependencias:** como mucho dos dependencias directas en go.mod y ninguna que sea de otro framework.
2. **Mensaje de fallo:** si dos listas no coinciden, el mensaje muestra la diferencia sin escribirla a mano.
3. **Tipos:** comparar valores de tipos distintos es un error de compilación, no un fallo al ejecutar.
4. **Sin escribir en el directorio del código fuente:** o, si escribe algo, que la ruta se pueda configurar (objetivo 5)
5. **Sin instalaciones:** solo hace falta Go para ejecutarlo en el contenedor y en la CI.
6. **Mantenida:** algún cambio en los dos últimos años.
7. **Complejidad ciclomática mínima en los tests**: las comprobaciones no añaden ramas al test

Las opciones las he buscado en la sección Testing de [awesome-go](https://github.com/avelino/awesome-go#testing) y en [pkg.go.dev](https://pkg.go.dev). De cada candidata he mirado su go.mod para ver las dependencias directas y la fecha de su ultimo commit.

## Framework y CLI
- **[`testing`](https://pkg.go.dev/testing) + `go test`** (framework y CLI): Viene con Go, así que no añaden dependencias ni hay que instalar nada. `testing` tiene subtests, preparación y limpieza y ejecución en paralelo. `go test` guarda la caché y los módulos en `GOMODCACHE` y las dos rutas se configuran con variables de entorno. No modifica `go.mod` ni `go.sum`. `go.sum` solo cambia al añadir una dependencia con `go get` y eso se hace en desarrollo.
- **[Ginkgo](https://github.com/onsi/ginkgo) v2.34.0** (framework BDD con CLI opcional): sus tests también se ejecutan con `go test` así que no obliga a instalar nada. Pero añade unas diez dependencias directas y no cumple el criterio 1
- **[gotestsum](https://github.com/gotestyourself/gotestsum) v1.13.0** (solo CLI): formatea la salida de `go test`. Hay que instalarlo en el contenedor y en la CI, así que no cumple el criterio 5, y lo que aporta no resuelve ningún problema de la HU001.

## Biblioteca de aserciones

`testing` no tiene aserciones: se compara a mano con `if` y `t.Errorf`. Las opciones:

- **Sin biblioteca** (`if` + `t.Errorf`): no añade dependencias y se detectan los distintos tipos al compilar. No cumple el criterio 2, porque cada mensaje de fallo se escribe a mano y no muestra la diferenca entre las listas.
- **[testify](https://github.com/stretchr/testify) v1.12.1**: añade dos dependencias, muestra la diferencia y se mantiene (último cambio en septiembre de 2026). No cumple el criterio 3: compara valorest `interface{}`, así que un error de tipo solo aparece al ejecutar.
- **[gotest.tools](https://github.com/gotestyourself/gotest.tools) v3.5.2**: añade dos dependencias, muestra la diferencia con go-cmp y se mantiene (septiembre de 2026). No cumple el criterio 3 por el mismo motivo que testify.
- **[Gomega](https://github.com/onsi/gomega) v1.44.0**: estilo BDD, pensada para Ginkgo. Añade cinco dependencias, entre ellas Ginkgo y protobuf, y tampoco comprueba los tipos al compilar. No cumple los criterios 1 y 3.
- **[matryer/is](https://github.com/matryer/is) v1.4.1**: no añade dependencias, pero no muestra la diferencia, no comprueba los tipos y su último cambio es de mayo de 2023. No cumple los criterios 2, 3 y 6.
- **[alecthomas/assert](https://github.com/alecthomas/assert) v2.11.0**: Añade dos pequeñas dependencias, muestra la diferencia y al ser genérica, comparar tipos distintos no compila. Su último cambio es de octubre de 2025. Cumple los 4 criterios que se le aplican: 1,2,3 y 6

El mismo test, roto a propósito (el sábado se espera un hueco en el horno y se obtiene en el teléfono), con los tres finalistas:

![Mensajes de fallo de las tres finalistas](..img/Objetivo-4/mensajes-fallo.png)

- **Sin biblioteca:** imprime las dos listas completas, y hay que encontrar a mano qué elemento cambia.
- **testify:** imprime las dos listas completas y depsues un `Diff` que señala `Horno` frente a `Telefono`, con los días como números (`Dia: 6`).
- **alecthomas/assert:** muestra solo la parte que cambia, con el día por su nombre. Omite los campos con valor cero y como en el domingo es time.Weekday(0), ese elemento aparece sin Dia. Lo tendré en cuenta en los tests de turnos de domingo
