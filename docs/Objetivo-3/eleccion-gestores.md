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



  
