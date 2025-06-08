package main

import (
	"log"

	pb "github.com/ocdsquad/golang-grpc/proto" // Adjust the import path to your actual project structure
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8084"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure()) // Deprecated, use WithTransportCredentials instead
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		panic(err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	log.Printf("Client connected to server on %s", port)

	names := &pb.NameList{
		Names: []string{"Alice", "Bob", "Charlie"},
	}

	// callSayHello(client)
	// callGreetServerStream(client, names)
	callGreetClientStream(client, names)
	// callGreetBidirectionalStream(client, names)
}
