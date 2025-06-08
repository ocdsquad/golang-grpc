package main

import (
	"context"

	pb "github.com/ocdsquad/golang-grpc/proto" // Adjust the import path to your actual project structure
)

func (s *server) Greet(ctx context.Context, req *pb.NoParam) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Message: "Hello, World!",
	}, nil
}
