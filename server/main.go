package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"main/database"
	pb "main/go_protocol_buffer"
	cruds "main/server/crud"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceServer
}

func (s *server) AnonymousSignUp(ctx context.Context, req *pb.AnonymousSignUpRequest) (*pb.AnonymousSignUpResponse, error) {
	db := database.GetDB()

	user, err := cruds.CreateUser(db, req.Name)
	if err != nil {
		return nil, err
	}

	accessToken, err := BuildAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	response := pb.AnonymousSignUpResponse{
		User: &pb.User{Id: uint64(user.ID), Name: user.Name}, AccessToken: accessToken,
	}

	return &response, nil
}

func (s *server) CreateShareGroup(ctx context.Context, req *pb.CreateShareGroupRequest) (
	*pb.CreateShareGroupResponse, error,
) {
	claims := Auth(ctx)
	userId := claims["ist"]
	db := database.GetDB()

	shareGroup, err := cruds.CreateShareGroup(db, req.DestLng, req.DestLat, req.MeetingTime)
	if err != nil {
		return nil, err
	}

	fmt.Println(userId)

	return shareGroup, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, &server{})
	fmt.Println("Server is running on port: 8080")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
