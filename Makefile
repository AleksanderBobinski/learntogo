.PHONY: all fmt lint test

all: fmt lint test client

fmt:
	go fmt ./cmd/client

lint:
	golangci-lint run ./cmd/client

test:
	go test -cover ./cmd/client

client:
	go build ./cmd/client
