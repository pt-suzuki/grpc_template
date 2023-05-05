.PHONY: generate.*
generate.proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. -I./proto ./proto/*.proto ./proto/*/*/*.proto

generate.wire:
	cd src/provider && \
	wire

.PHONY: run
run:
	go run cmd/main.go
