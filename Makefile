.PHONY: build test lint lint-fix

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v -race -cover ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
