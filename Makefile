start:
	go run cmd/main.go cmd/wire_gen.go

wire:
	wire ./cmd

protoc:
	cd ./internal/infra/grpc/ && protoc --go_out=. --go-grpc_out=. protofiles/order.proto

up:
	docker-compose up -d

down:
	docker-compose down