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
		viewUser.Lat = user.Lat
	}
	if user.Lon != nil {
		viewUser.Lon = user.Lon
	}
	if user.PositionAt != nil {
		positionAtStr := (*user.PositionAt).String()
		viewUser.PositionAt = &positionAtStr
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
		DestLon:     shareGroup.DestLon,
		DestLat:     shareGroup.DestLat,
		LinkKey:     shareGroup.LinkKey,
		MeetingTime: shareGroup.MeetingTime,
		InviteUrl:   fmt.Sprintf("mesh://invite/%s", shareGroup.LinkKey),
		Address:     shareGroup.Address,
		Users:       UserListMapper(shareGroup.Users),
		AdminUser:   UserMapper(&shareGroup.AdminUser),
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

func GetShareGroupLinkKeyResponseMapper(shareGroup *database.ShareGroup) *pb.GetShareGroupByLinkKeyResponse {
	return &pb.GetShareGroupByLinkKeyResponse{
		ShareGroup: ShareGroupMapper(shareGroup),
	}
}

func GetCurrentShareGroupResponseMapper(shareGroup *database.ShareGroup) *pb.GetCurrentShareGroupResponse {
	return &pb.GetCurrentShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup),
	}
}

func UpdatePositionGroupResponseMapper(user *database.User) *pb.UpdatePositionResponse {
	return &pb.UpdatePositionResponse{
		User: UserMapper(user),
	}
}

func GetCurrentUserResponseMapper(user *database.User) *pb.GetCurrentUserResponse {
	return &pb.GetCurrentUserResponse{
		User: UserMapper(user),
	}
}

func LeaveShareGroupResponseMapper(user *database.User) *pb.LeaveShareGroupResponse {
	return &pb.LeaveShareGroupResponse{
		User: UserMapper(user),
	}
}
