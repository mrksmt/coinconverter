.PHONY: lint test all

build:
	go build -o coinconv  cmd/console_app/main.go

all: lint test build

lint:
	 golangci-lint run

test:
	go test ./...

run_code:
	go run cmd/console_app/main.go 123.45 USD BTC

run_error:
	./coinconv 123.45 USD ччч

run:
	./coinconv 123.45 USD BTC




