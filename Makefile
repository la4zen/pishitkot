dev:
	WORKDIR=$PWD go run ./cmd/app/main.go

build:
	go build ./cmd/app/main.go app

test:
	go test ./...