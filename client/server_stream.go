package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ocdsquad/golang-grpc/proto"
)

func callGreetServerStream(client pb.GreetServiceClient, names *pb.NameList) {

	log.Printf("Streaming start")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.GreetServerStream(ctx, names)

	if err != nil {
		log.Fatalf("Error calling GreetServerStream: %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving message from GreetServerStream: %v", err)
			break
		}
		log.Printf("Received message: %s", message.Message)
	}
	log.Printf("Streaming end")
	log.Printf("GreetServerStream completed successfully")

}
