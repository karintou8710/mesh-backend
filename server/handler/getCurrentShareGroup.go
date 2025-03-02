package handler

import (
	"context"
	"fmt"
	"main/database"
	pb "main/go_protocol_buffer"
	"main/server/auth"
	"main/server/crud"
	"main/server/view"
)

func (s *Server) GetCurrentShareGroup(ctx context.Context, req *pb.GetCurrentShareGroupRequest) (
	*pb.GetCurrentShareGroupResponse, error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	shareGroup := crud.GetShareGroupByLinkKey(db, user.ShareGroup.LinkKey)
	if shareGroup == nil {
		return nil, fmt.Errorf("error: the linkKey is invalid")
	}

	return view.GetCurrentShareGroupResponseMapper(shareGroup), nil
}
