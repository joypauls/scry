build:
	go build

build-small:
	go build -ldflags="-s -w"

run:
	go run main.go
