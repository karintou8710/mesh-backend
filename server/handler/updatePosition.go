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

func (s *Server) UpdatePosition(ctx context.Context, req *pb.UpdatePositionRequest) (
	*pb.UpdatePositionResponse, error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	user, err := crud.UpdatePosition(db, user, req.Lat, req.Lon)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, fmt.Errorf("error: failed to update position")
	}

	return view.UpdatePositionGroupResponseMapper(user), nil
}
