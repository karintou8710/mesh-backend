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

func (s *Server) ArriveDest(ctx context.Context, req *pb.ArriveDestRequest) (
	*pb.ArriveDestResponse, error,
) {
	db := database.GetDB()

	user := auth.Auth(ctx)
	if user == nil {
		return nil, fmt.Errorf("error: need to authenticate")
	}

	user, err := crud.UpdateIsArrived(db, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return view.ArriveDestResponseMapper(user), nil
}
