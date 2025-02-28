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

	shareGroup, err := cruds.CreateShareGroup(db, req.DestLon, req.DestLat, req.MeetingTime, req.Address, user.ID)
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

func (s *server) JoinShareGroup(ctx context.Context, req *pb.JoinShareGroupRequest) (
	*pb.JoinShareGroupResponse, error,
) {
	db := database.GetDB()

	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	shareGroup := cruds.GetShareGroupByLinkKey(db, req.LinkKey)
	if shareGroup == nil {
		return nil, fmt.Errorf("error: the linkKey is invalid")
	}

	shareGroup, err := cruds.JoinShareGroup(db, shareGroup, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return JoinShareGroupResponseMapper(shareGroup), nil
}

func (s *server) GetShareGroupByLinkKey(ctx context.Context, req *pb.GetShareGroupByLinkKeyRequest) (
	*pb.GetShareGroupByLinkKeyResponse, error,
) {
	db := database.GetDB()

	shareGroup := cruds.GetShareGroupByLinkKey(db, req.LinkKey)
	if shareGroup == nil {
		return nil, fmt.Errorf("error: the linkKey is invalid")
	}

	return GetShareGroupLinkKeyResponseMapper(shareGroup), nil
}

func (s *server) GetCurrentShareGroup(ctx context.Context, req *pb.GetCurrentShareGroupRequest) (
	*pb.GetCurrentShareGroupResponse, error,
) {
	db := database.GetDB()

	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	shareGroup := cruds.GetShareGroupByLinkKey(db, user.ShareGroup.LinkKey)
	if shareGroup == nil {
		return nil, fmt.Errorf("error: the linkKey is invalid")
	}

	return GetCurrentShareGroupResponseMapper(shareGroup), nil
}

func (s *server) UpdatePosition(ctx context.Context, req *pb.UpdatePositionRequest) (
	*pb.UpdatePositionResponse, error,
) {
	db := database.GetDB()

	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	user, err := cruds.UpdatePosition(db, user, req.Lat, req.Lon)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, fmt.Errorf("error: failed to update position")
	}

	return UpdatePositionGroupResponseMapper(user), nil
}

func (s *server) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (
	*pb.GetCurrentUserResponse, error,
) {
	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	return GetCurrentUserResponseMapper(user), nil
}

func (s *server) LeaveShareGroup(ctx context.Context, req *pb.LeaveShareGroupRequest) (
	*pb.LeaveShareGroupResponse, error,
) {
	db := database.GetDB()

	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	user, err := cruds.LeaveShareGroup(db, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return LeaveShareGroupResponseMapper(user), nil
}

func (s *server) ArriveDest(ctx context.Context, req *pb.ArriveDestRequest) (
	*pb.ArriveDestResponse, error,
) {
	db := database.GetDB()

	user := Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	user, err := cruds.UpdateIsArrived(db, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ArriveDestResponseMapper(user), nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
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
