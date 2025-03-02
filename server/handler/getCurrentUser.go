package handler

import (
	"context"
	"fmt"

	pb "main/go_protocol_buffer"
	"main/server/auth"
	"main/server/view"
)

func (s *Server) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (
	*pb.GetCurrentUserResponse, error,
) {
	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	return view.GetCurrentUserResponseMapper(user), nil
}
