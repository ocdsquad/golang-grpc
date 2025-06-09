package main

import (
	"io"
	"log"

	pb "github.com/ocdsquad/golang-grpc/proto"
)

func (s *server) GreetBidirectionalStream(stream pb.GreetService_GreetBidirectionalStreamServer) error {
	log.Printf("GreetBidirectionalStream started")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil // End of stream
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}
		log.Printf("Received name: %s", req.Name)

		response := &pb.GreetResponse{
			Message: "Hello, " + req.Name,
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}
	}
}
