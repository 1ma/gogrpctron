package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/1ma/gogrpctron/internal/grpc"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func (s *userService) CreateUser(ctx context.Context, in *emptypb.Empty) (*pb.UserData, error) {
	return &pb.UserData{Id: 1, UserName: "Christiana Hessel", CreatedAt: timestamppb.Now()}, nil
}

func (s *userService) ListUsers(ctx context.Context, in *emptypb.Empty) (*pb.UserList, error) {
	return &pb.UserList{Users: []*pb.UserData{{Id: 1, UserName: "Christiana Hessel", CreatedAt: timestamppb.Now()}}}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userService{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
