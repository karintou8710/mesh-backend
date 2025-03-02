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

func (s *Server) JoinShareGroup(ctx context.Context, req *pb.JoinShareGroupRequest) (
	*pb.JoinShareGroupResponse, error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	shareGroup := crud.GetShareGroupByLinkKey(db, req.LinkKey)
	if shareGroup == nil {
		return nil, fmt.Errorf("error: the linkKey is invalid")
	}

	shareGroup, err := crud.JoinShareGroup(db, shareGroup, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return view.JoinShareGroupResponseMapper(shareGroup), nil
}
