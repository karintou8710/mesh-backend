package view

import (
	"fmt"
	"log"
	"main/database"
	pb "main/go_protocol_buffer"
	"main/server/utils"
	"slices"
)

func UserWithoutPositionMapper(user *database.User) *pb.User {
	if user == nil {
		return nil
	}

	viewUser := &pb.User{
		Id:         uint64(user.ID),
		Name:       user.Name,
		ShareGroup: ShareGroupMapper(user.ShareGroup, []uint{}),
		IsArrived:  user.IsArrived,
		IconID:     user.IconID,
	}
	if user.ShortMessage != nil {
		viewUser.ShortMessage = user.ShortMessage
	}
	if user.ShareGroupID != nil {
		viewUser.ShareGroupId = user.ShareGroupID
	}
	viewUser.ShareGroup = ShareGroupMapper(user.ShareGroup, []uint{})

	return viewUser
}

func UserMapper(user *database.User) *pb.User {
	if user == nil {
		return nil
	}

	viewUser := &pb.User{
		Id:         uint64(user.ID),
		Name:       user.Name,
		ShareGroup: ShareGroupMapper(user.ShareGroup, []uint{}),
		IsArrived:  user.IsArrived,
		IconID:     user.IconID,
	}
	if user.ShortMessage != nil {
		viewUser.ShortMessage = user.ShortMessage
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
		viewUser.ShareGroupId = user.ShareGroupID
	}
	viewUser.ShareGroup = ShareGroupMapper(user.ShareGroup, []uint{})

	return viewUser
}

func UserListMapper(users []*database.User) []*pb.User {
	viewUsers := []*pb.User{}

	for _, user := range users {
		viewUsers = append(viewUsers, UserMapper(user))
	}

	return viewUsers
}

func UserWithoutPositionListMapper(users []*database.User, includeIds []uint) []*pb.User {
	viewUsers := []*pb.User{}

	for _, user := range users {
		if slices.Contains(includeIds, user.ID) {
			viewUsers = append(viewUsers, UserMapper(user))
		} else {
			viewUsers = append(viewUsers, UserWithoutPositionMapper(user))
		}
	}

	return viewUsers
}

func ShareGroupMapper(shareGroup *database.ShareGroup, includeIds []uint) *pb.ShareGroup {
	if shareGroup == nil {
		return nil
	}

	viewShareGroup := &pb.ShareGroup{
		Id:                       uint64(shareGroup.ID),
		DestLon:                  shareGroup.DestLon,
		DestLat:                  shareGroup.DestLat,
		LinkKey:                  shareGroup.LinkKey,
		MeetingTime:              shareGroup.MeetingTime,
		InviteUrl:                fmt.Sprintf("https://redirect-project-ashy.vercel.app/invite/%s", shareGroup.LinkKey),
		Address:                  shareGroup.Address,
		SharingLocationStartTime: shareGroup.SharingLocationStartTime,
		IsSharingLocation:        true,
	}

	if shareGroup.SharingLocationStartTime != nil {
		sharingLocationStartTime, err := utils.ParseISO8601(*shareGroup.SharingLocationStartTime)
		if err == nil {
			viewShareGroup.IsSharingLocation = utils.IsPastTime(sharingLocationStartTime)
		} else {
			log.Println(err)
		}
	}

	if viewShareGroup.IsSharingLocation {
		viewShareGroup.Users = UserListMapper(shareGroup.Users)
		viewShareGroup.AdminUser = UserMapper(&shareGroup.AdminUser)
	} else {
		viewShareGroup.Users = UserWithoutPositionListMapper(shareGroup.Users, includeIds)
		viewShareGroup.AdminUser = UserWithoutPositionMapper(&shareGroup.AdminUser)
	}

	return viewShareGroup
}

func AnonymousSignUpResponseMapper(user *database.User, accessToken string) *pb.AnonymousSignUpResponse {
	return &pb.AnonymousSignUpResponse{
		User: UserMapper(user), AccessToken: accessToken,
	}
}

func CreateShareGroupResponseMapper(shareGroup *database.ShareGroup) *pb.CreateShareGroupResponse {
	return &pb.CreateShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup, []uint{}),
	}
}

func JoinShareGroupResponseMapper(shareGroup *database.ShareGroup) *pb.JoinShareGroupResponse {
	return &pb.JoinShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup, []uint{}),
	}
}

func GetShareGroupLinkKeyResponseMapper(shareGroup *database.ShareGroup) *pb.GetShareGroupByLinkKeyResponse {
	return &pb.GetShareGroupByLinkKeyResponse{
		ShareGroup: ShareGroupMapper(shareGroup, []uint{}),
	}
}

func GetCurrentShareGroupResponseMapper(shareGroup *database.ShareGroup, currentUserID uint) *pb.GetCurrentShareGroupResponse {
	return &pb.GetCurrentShareGroupResponse{
		ShareGroup: ShareGroupMapper(shareGroup, []uint{currentUserID}),
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

func LeaveShareGroupResponseMapper() *pb.LeaveShareGroupResponse {
	return &pb.LeaveShareGroupResponse{}
}

func ArriveDestResponseMapper(user *database.User) *pb.ArriveDestResponse {
	return &pb.ArriveDestResponse{
		User: UserMapper(user),
	}
}

func UpdateShortMessageMapper(user *database.User) *pb.UpdateShortMessageResponse {
	return &pb.UpdateShortMessageResponse{
		User: UserMapper(user),
	}
}
