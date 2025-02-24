package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"main/database"
	pb "main/go_protocol_buffer"

	"github.com/golang-jwt/jwt/v5"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 365).Unix(),
	})

	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	response := pb.AnonymousSignUpResponse{
		User: &pb.User{Id: uint64(user.ID), Name: user.Name}, AccessToken: accessToken,
	}

	return &response, nil
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
