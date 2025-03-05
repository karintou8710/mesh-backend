package handler

import (
	"context"
	"fmt"
	"log"
	"main/database"
	pb "main/go_protocol_buffer"
	"main/server/auth"
	"main/server/crud"
	"main/server/view"
)

func (s *Server) LeaveShareGroup(ctx context.Context, req *pb.LeaveShareGroupRequest) (
	*pb.LeaveShareGroupResponse, error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	// ShareGroupに参加していない場合
	if user.ShareGroup == nil {
		return view.LeaveShareGroupResponseMapper(), nil
	}

	if user.ID == uint(user.ShareGroup.AdminUserID) {
		err := crud.DeleteShareGroupByLinkKey(db, user.ShareGroup.LinkKey)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else {
		err := crud.LeaveShareGroup(db, user)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return view.LeaveShareGroupResponseMapper(), nil
}
