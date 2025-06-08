package main

import (
	"io"
	"log"

	pb "github.com/ocdsquad/golang-grpc/proto" // Adjust the import path to your actual project structure
)

func (s *server) GreetClientStream(stream pb.GreetService_GreetClientStreamServer) error {
	log.Printf("GreetClientStream started")

	var messages []string

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.MessagesList{Messages: messages})
			}
			log.Printf("Error receiving message: %v", err)
			return err
		}
		log.Printf("Received name: %s", req.Name)
		messages = append(messages, "Hello", req.Name)
	}

}
