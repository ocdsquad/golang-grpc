package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ocdsquad/golang-grpc/proto"
)

func callGreetClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming start")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.GreetClientStream(ctx)
	if err != nil {
		log.Fatalf("Error calling GreetClientStream: %v", err)
	}

	for _, name := range names.Names {
		log.Printf("Sending name: %s", name)
		if err := stream.Send(&pb.GreetRequest{Name: name}); err != nil {
			log.Fatalf("Error sending name: %v", err)
		}
		time.Sleep(1 * time.Second) // Simulate some delay between sends
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response from GreetClientStream: %v", err)
	}
	log.Printf("Received messages: %v", resp.Messages)
	log.Printf("Streaming end")
}
