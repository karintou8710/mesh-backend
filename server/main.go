package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"main/database"
	pb "main/go_protocol_buffer"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) AnonymousSignUp(ctx context.Context, req *pb.AnonymousSignUpRequest) (*pb.AnonymousSignUpResponse, error) {
	db := database.GetDB()

	user := database.User{Name: req.Name}
	if res := db.Create(&user); res.Error != nil {
		log.Fatalf("Error: %v", res.Error)
		return nil, res.Error
	}

	return &pb.AnonymousSignUpResponse{User: &pb.User{Id: uint64(user.ID), Name: user.Name}}, nil
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
