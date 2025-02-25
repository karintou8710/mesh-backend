package main

import (
	"fmt"
	"main/database"
	pb "main/go_protocol_buffer"
)

func UserMapper(user *database.User) *pb.User {
	if user == nil {
		return nil
	}

	viewUser := &pb.User{
		Id:         uint64(user.ID),
		Name:       user.Name,
		ShareGroup: ShareGroupMapper(user.ShareGroup),
	}
	if user.Lat != nil {
		viewUser.Lat = *user.Lat
	}
	if user.Lon != nil {
		viewUser.Lon = *user.Lon
	}
	if user.PositionAt != nil {
		viewUser.PositionAt = (*user.PositionAt).String()
	}
	if user.ShareGroupID != nil {
		viewUser.ShareGroupId = *user.ShareGroupID
	}
	viewUser.ShareGroup = ShareGroupMapper(user.ShareGroup)

	return viewUser
}

func UserListMapper(users []*database.User) []*pb.User {
	viewUsers := []*pb.User{}

	for _, user := range users {
		viewUsers = append(viewUsers, UserMapper(user))
	}

	return viewUsers
}

func ShareGroupMapper(shareGroup *database.ShareGroup) *pb.ShareGroup {
	if shareGroup == nil {
		return nil
	}

	return &pb.ShareGroup{
		Id:          uint64(shareGroup.ID),
		DestLng:     shareGroup.DestLng,
		DestLat:     shareGroup.DestLat,
		LinkKey:     shareGroup.LinkKey,
		MeetingTime: shareGroup.MeetingTime,
		InviteUrl:   fmt.Sprintf("mesh://invite/%s", shareGroup.LinkKey),
		Users:       UserListMapper(shareGroup.Users),
	}
}

func AnonymousSignUpResponseMapper(user *database.User, accessToken string) *pb.AnonymousSignUpResponse {
	return &pb.AnonymousSignUpResponse{
		User: UserMapper(user), AccessToken: accessToken,
	}
}

func CreateShareGroupResponseMapper(shareGroup *database.ShareGroup) *pb.CreateShareGroupResponse {
	return &pb.CreateShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup),
	}
}

func JoinShareGroupResponseMapper(shareGroup *database.ShareGroup) *pb.JoinShareGroupResponse {
	return &pb.JoinShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup),
	}
}

func GetCurrentShareGroupResponseMapper(shareGroup *database.ShareGroup) *pb.GetCurrentShareGroupResponse {
	return &pb.GetCurrentShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup),
	}
}

func UpdatePositionGroupResponseMapper(user *database.User) *pb.UpdatePositionGroupResponse {
	return &pb.UpdatePositionGroupResponse{
		User: UserMapper(user),
	}
}
