BIN_NAME=sbit-emitter

dep:
	@echo "> fetching dependencies..."
	@go mod tidy

install: dep
	@echo "> building binary..."
	@CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ./${BIN_NAME} main.go


run: dep install
	@echo "> running application..."
	@./${BIN_NAME} serve

protoc-gen:
	@echo "> generating protobuf files..."
	@cd ./shared/proto/emitter && protoc *.proto -I. --go_out=../../../infrastructure/server/grpc/proto/emitter  --go-grpc_out=require_unimplemented_servers=false:../../../infrastructure/server/grpc/proto/emitter

protoc-gen-clean:
	@echo "> cleaning protobuf files..."
	#@cd ./infrastructure/server/grpc/proto/validator && rm -rf *.pb.go
	@cd ./infrastructure/server/grpc/proto/emitter && rm -rf *.pb.go

eal.PHONY: dep install run protoc-gen protoc-gen-clean setup