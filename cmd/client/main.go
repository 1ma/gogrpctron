package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/1ma/gogrpctron/internal/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	response, err := c.ListUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %v", response)
}
