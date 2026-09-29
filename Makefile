.PHONY: build test test-coverage lint lint-fix

COVERAGE_MIN = 80

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v -race -cover ./...

# пакет main (cmd/) не тестируется, поэтому в покрытие не входит
test-coverage:
	go test -coverprofile=coverage.out $$(go list ./... | grep -v /cmd/)
	go tool cover -func=coverage.out | awk '/^total:/ { print; if ($$3 + 0 < $(COVERAGE_MIN)) exit 1 }'

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
