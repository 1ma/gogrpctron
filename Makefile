.PHONY: dev clean protoc

dev: protoc
	go run cmd/server/main.go

clean:
	rm -f internal/grpc/*.pb.go

protoc:
	protoc --proto_path=/usr/local/include --proto_path=internal/grpc --go_out=. --go-grpc_out=. internal/grpc/user_service.proto
