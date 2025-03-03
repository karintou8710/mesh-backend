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

func (s *Server) UpdateShortMessage(ctx context.Context, req *pb.UpdateShortMessageRequest) (
	*pb.UpdateShortMessageResponse,
	error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	user, err := crud.UpdateShortMessage(db, user, &req.ShortMessage)
	if err != nil {
		return nil, err
	}

	return view.UpdateShortMessageMapper(user), nil
}
