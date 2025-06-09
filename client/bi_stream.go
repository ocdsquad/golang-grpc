package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ocdsquad/golang-grpc/proto"
)

func callGreetBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional streaming start")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.GreetBidirectionalStream(ctx)
	if err != nil {
		log.Fatalf("Error calling GreetBidirectionalStream: %v", err)
	}

	waitch := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error receiving message: %v", err)

				return
			}
			log.Printf("Received message: %s", message.Message)
		}
		close(waitch)
	}()

	// Send names to the server
	for _, name := range names.Names {
		if err := stream.Send(&pb.GreetRequest{Name: name}); err != nil {
			log.Fatalf("Error sending name: %v", err)
		}
		log.Printf("Sent name: %s", name)
		// Simulate some delay between sends
		time.Sleep(1 * time.Second) //simulate a delay
	}

	// Close the stream to indicate no more messages will be sent
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing send stream: %v", err)
	}

	// <- wait for the server to finish processing
	<-waitch

	log.Printf("Bidirectional streaming end")
}
