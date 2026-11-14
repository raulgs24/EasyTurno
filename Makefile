.PHONY: install check

# Descarga los módulos declarados en el go.mod. Ahora no hay ninguno
# (solo utiliza biblioteca estándar), así que no descarga nada;
# queda preparada para cuando llegue la primera dependencia.

install:
	go mod download 

check:
	go build ./...
