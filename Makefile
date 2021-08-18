build:
	go build -o ./bin/scry

build-small:
	go build -ldflags="-s -w" -o ./bin/scry

run:
	go run main.go

dev:
	sh ./dev-setup.sh

clean:
	go mod tidy
