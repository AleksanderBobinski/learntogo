.PHONY: all fmt lint test

all: fmt lint static test client server

proto/message_grpc.pb.go proto/message.pb.go: proto/message.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/message.proto

fmt:
	go fmt ./pkg/connection/
	go fmt ./cmd/client/
	go fmt ./cmd/server/

lint: proto/message_grpc.pb.go proto/message.pb.go
	golangci-lint run ./pkg/connection/
	golangci-lint run ./cmd/client/
	golangci-lint run ./cmd/server/

static: proto/message_grpc.pb.go proto/message.pb.go
	 staticcheck ./pkg/connection/
	 staticcheck ./cmd/client/
	 staticcheck ./cmd/server/

test: client server
	go test -cover ./cmd/client
	go test -cover ./cmd/server

client: proto/message_grpc.pb.go proto/message.pb.go ./cmd/client/client.go ./pkg/connection/connection.go
	go build ./cmd/client

server: proto/message_grpc.pb.go proto/message.pb.go ./cmd/client/client.go ./pkg/connection/connection.go
	go build ./cmd/server
