proto:
	protoc --go-grpc_out=. ./pkg/pb/proto/*.proto

	

	
server:
	go run cmd/main.go	

wire: ## Generate wire_gen.go
	cd pkg/di && wire	