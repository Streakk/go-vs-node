package main

import (
	"context"
	"fmt"
	profile "grpc-rest-test/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	profile.UnimplementedProfileServiceServer
}

func (s *server) GetUser(ctx context.Context, req *profile.GetUserRequest) (*profile.User, error) {
	return &profile.User{
		UserID:   req.UserID,
		Username: "JohnDoe",
		Email:    "john.doe@example.com",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	profile.RegisterProfileServiceServer(s, &server{})
	fmt.Println("Listning on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
