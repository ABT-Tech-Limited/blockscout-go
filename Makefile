.PHONY: test lint build

build:
	go build ./...

test:
	@GOMAXPROCS=1 go test -p=1 ./... -v

lint:
	golangci-lint run ./...

tidy:
	go mod tidy
