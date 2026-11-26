install:
	go mod download 

check:
	go build ./...

test:
	go test ./...