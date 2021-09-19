.PHONY: dev mysql clean

dev:
	protoc --proto_path=/usr/local/include --proto_path=internal/grpc --go_out=. --go-grpc_out=. internal/grpc/user_service.proto
	docker-compose run --rm starter
	go run cmd/server/main.go

mysql:
	docker-compose exec mysql mysql -D gogrpctron

clean:
	rm -f internal/grpc/*.pb.go
