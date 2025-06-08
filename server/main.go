package main

import (
	"log"
	"net"

	pb "github.com/ocdsquad/golang-grpc/proto" // Adjust the import path to your actual project structure
	"google.golang.org/grpc"
)

const (
	port = ":8084"
)

type server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})
	log.Printf("Server is listening on %s", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// Register your gRPC services here

}
