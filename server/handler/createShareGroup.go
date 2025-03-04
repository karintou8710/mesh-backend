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

func (s *Server) CreateShareGroup(ctx context.Context, req *pb.CreateShareGroupRequest) (
	*pb.CreateShareGroupResponse, error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	shareGroup, err := crud.CreateShareGroup(db, req.DestLon, req.DestLat, req.MeetingTime, req.Address, user.ID, req.SharingLocationStartTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	shareGroup, err = crud.JoinShareGroup(db, shareGroup, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return view.CreateShareGroupResponseMapper(shareGroup), nil
}
