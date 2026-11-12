#!/bin/sh
# Mediciones para docs/gestor-dependencias.md y docs/gestor-tareas.md
# Uso, desde la raíz del repo:  sh scripts/mediciones.sh <medida>
#   gccgo | toolchain | gomod | presentes | tamano | gosum | check-cmd | phony | tiempos
# Necesita Docker. Las pruebas se hacen en una copia temporal, el repo no se toca.

BASE=golang:1.27-alpine
REPO=$(pwd)

copia() {
  tmp=$(mktemp -d)
  cp "$REPO/go.mod" "$tmp/" && cp -r "$REPO/internal" "$tmp/"
  echo "$tmp"
}

case "$1" in

gccgo)  # Compilar el proyecto con gccgo (imagen oficial de GCC)
  docker run --rm -v "$REPO":/src -w /src gcc:15 \
    sh -c 'gccgo --version | head -1; go version; go build ./... && echo COMPILA'
  ;;

toolchain)  # Qué pasa con una versión de Go anterior a la del go.mod
  docker run --rm -v "$REPO":/src -w /src -e GOTOOLCHAIN=local golang:1.26-alpine \
    sh -c 'go version; go build ./...'
  docker run --rm -v "$REPO":/src -w /src golang:1.26-alpine \
    sh -c 'time go build ./... && go version'
  docker run --rm $BASE go version
  ;;

gomod)  # Dependencias reales del proyecto y verificación
  go mod download; go mod verify; go env GOPROXY GOSUMDB GOTOOLCHAIN
  ls go.sum 2>/dev/null || echo "No existe go.sum"
  ;;

presentes)  # ¿Viene ya instalado cada gestor en las imágenes?
  for img in golang:1.27-alpine golang:1.27-bookworm; do
    echo "== $img"
    docker run --rm "$img" sh -c \
      'for t in make task just mage; do printf "%s: " $t; command -v $t || echo "no está"; done'
  done
  echo "ubuntu-latest: mirar la lista de software de actions/runner-images"
  ;;

tamano)  # KiB que añade cada uno a la imagen ligera
  docker pull -q $BASE >/dev/null
  construye() { printf 'FROM %s\nRUN %s\n' "$BASE" "$2" | docker build -q -t "medida-$1" - >/dev/null; }
  construye make  "apk add --no-cache make"
  construye just  "apk add --no-cache just"
  construye task  "wget -qO- https://taskfile.dev/install.sh | sh -s -- -d -b /usr/local/bin"
  construye mage  "go install github.com/magefile/mage@latest && go clean -cache -modcache"
  base=$(docker image inspect -f '{{.Size}}' $BASE)
  for t in make just task mage; do
    s=$(docker image inspect -f '{{.Size}}' "medida-$t")
    echo "$t: +$(( (s - base) / 1024 )) KiB"
  done
  echo "goyek: no se instala, es una biblioteca (ver gosum)"
  ;;

gosum)  # Lo que entra en go.mod/go.sum si se declara como herramienta o biblioteca
  for h in "-tool github.com/go-task/task/v3/cmd/task@latest" \
           "-tool github.com/magefile/mage@latest" \
           "github.com/goyek/goyek/v2@latest"; do
    tmp=$(copia)
    (cd "$tmp" && go get $h >/dev/null 2>&1; echo "== go get $h"; wc -l go.mod go.sum 2>&1 | tail -n +1)
    rm -rf "$tmp"
  done
  ;;

check-cmd)  # Qué detecta cada orden candidata para "check"
  tmp=$(copia); cd "$tmp"; f=internal/domain/trabajador.go; cp $f /tmp/orig.go
  prueba() {
    echo "== $1"
    for c in "gofmt -e -l ." "go build ./..." "go vet ./..."; do
      $c >/dev/null 2>&1; echo "   $c -> salida $?"
    done
    cp /tmp/orig.go $f; rm -f internal/domain/z.go
  }
  prueba "Sin fallos"
  echo 'func roto( {' >> $f;                                   prueba "Error de sintaxis"
  printf 'func f() int { return "texto" }\n' >> $f;           prueba "Error de tipos"
  printf 'package domain\nimport "fmt"\nfunc g() { fmt.Printf("%%d", "x") }\n' > internal/domain/z.go
                                                              prueba "Printf mal usado (compila)"
  rm -rf "$tmp"
  ;;

phony)  # Por qué hace falta .PHONY
  tmp=$(copia); cd "$tmp"
  printf 'check:\n\tgo build ./...\n' > Makefile
  touch check
  echo "== Sin .PHONY y con un fichero llamado check:"; make check
  printf '.PHONY: check\ncheck:\n\tgo build ./...\n' > Makefile
  echo "== Con .PHONY:"; make check
  echo "== Reglas implícitas que trae make:"; make -p -f /dev/null 2>/dev/null | grep -c '^%'
  rm -rf "$tmp"
  ;;

tiempos)  # Tiempo de check dentro de la imagen ligera (hace falta haber lanzado "tamano")
  tmp=$(copia)
  printf '.PHONY: check\ncheck:\n\tgo build ./...\n' > "$tmp/Makefile"
  printf 'check:\n\tgo build ./...\n' > "$tmp/justfile"
  printf 'version: "3"\ntasks:\n  check:\n    cmds:\n      - go build ./...\n' > "$tmp/Taskfile.yml"
  printf '//go:build mage\n\npackage main\n\nimport "os/exec"\n\nfunc Check() error { return exec.Command("go", "build", "./...").Run() }\n' > "$tmp/magefile.go"
  for t in make just task mage; do
    echo "== $t"
    docker run --rm -v "$tmp":/src -w /src "medida-$t" \
      sh -c "apk add -q hyperfine && hyperfine -N --warmup 3 '$t check'"
  done
  rm -rf "$tmp"
  ;;

*) echo "Uso: sh scripts/mediciones.sh gccgo|toolchain|gomod|presentes|tamano|gosum|check-cmd|phony|tiempos" ;;
esac
