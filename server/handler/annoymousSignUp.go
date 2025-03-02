package handler

import (
	"context"
	"main/database"
	pb "main/go_protocol_buffer"
	"main/server/auth"
	"main/server/crud"
	"main/server/view"
)

func (s *Server) AnonymousSignUp(ctx context.Context, req *pb.AnonymousSignUpRequest) (
	*pb.AnonymousSignUpResponse,
	error,
) {
	db := database.GetDB()

	user, err := crud.CreateUser(db, req.Name)
	if err != nil {
		return nil, err
	}

	accessToken, err := auth.BuildAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return view.AnonymousSignUpResponseMapper(user, accessToken), nil
}
