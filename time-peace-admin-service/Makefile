proto:
	protoc --go-grpc_out=. ./pkg/pb/proto/*.proto
	


server:
	go run cmd/main.go

wire:
	cd pkg/di && wire
