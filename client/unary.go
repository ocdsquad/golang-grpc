package main

import (
	"context"
	"log"

	pb "github.com/ocdsquad/golang-grpc/proto" // Adjust the import path to your actual project structure
)

// TODO: add implementation for server side streaming, client streaming and bidirectional streaming

func callSayHello(client pb.GreetServiceClient) {
	// Create a request with no parameters

	ctx, cancel := context.WithCancel(context.Background())
	// Set a timeout for the request
	// ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	// Note: The context.WithCancel function is used to create a context that can be cancelled.
	// This is useful for setting a timeout or for cancelling the request if needed.
	defer cancel()
	req := &pb.NoParam{}

	// Call the Greet method
	resp, err := client.Greet(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Greet: %v", err)
	}

	// Print the response message
	log.Printf("Response from Greet: %s", resp.Message)
}
