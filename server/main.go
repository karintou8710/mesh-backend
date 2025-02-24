package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "main/go_protocol_buffer"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) AnonymousSignUp(ctx context.Context, req *pb.AnonymousSignUpRequest) (*pb.AnonymousSignUpResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.AnonymousSignUpResponse{}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &server{})

	fmt.Println("Server is running on port: 8080")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
