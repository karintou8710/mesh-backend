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

	return AnonymousSignUpResponseMapper(user, accessToken), nil
}

func (s *server) CreateShareGroup(ctx context.Context, req *pb.CreateShareGroupRequest) (
	*pb.CreateShareGroupResponse, error,
) {
	db := database.GetDB()

	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	shareGroup, err := cruds.CreateShareGroup(db, req.DestLng, req.DestLat, req.MeetingTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	shareGroup, err = cruds.JoinShareGroup(db, shareGroup, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return CreateShareGroupResponseMapper(shareGroup), nil
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
