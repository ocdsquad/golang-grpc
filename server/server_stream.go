package main

import (
	"log"
	"time"

	pb "github.com/ocdsquad/golang-grpc/proto" // Adjust the import path to your actual project structure
)

func (s *server) GreetServerStream(req *pb.NameList, stream pb.GreetService_GreetServerStreamServer) error {
	log.Printf("Received request with names: %v", req.Names)

	for _, name := range req.Names {
		// Simulate some processing
		response := &pb.GreetResponse{
			Message: "Hello, " + name + "!",
		}
		if err := stream.Send(response); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
		log.Printf("Sent response: %s", response.Message)
	}
	return nil
}
