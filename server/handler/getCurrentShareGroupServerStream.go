package handler

import (
	"fmt"
	"log"
	"main/database"
	pb "main/go_protocol_buffer"
	"main/server/auth"
	"main/server/crud"
	"main/server/view"
	"time"
)

func (s *Server) GetCurrentShareGroupServerStream(
	req *pb.GetCurrentShareGroupRequest,
	stream pb.Service_GetCurrentShareGroupServerStreamServer,
) error {
	db := database.GetDB()

	user := auth.Auth(stream.Context())
	if user == nil {
		return fmt.Errorf("error: need to authenticate")
	}

	for {
		if user.ShareGroup == nil {
			if err := stream.Send(view.GetCurrentShareGroupResponseMapper(nil, user.ID)); err != nil {
				log.Println(err)
				return err
			}
		} else {
			shareGroup := crud.GetShareGroupByLinkKey(db, user.ShareGroup.LinkKey)
			if shareGroup == nil {
				return fmt.Errorf("error: the linkKey is invalid")
			}

			if err := stream.Send(view.GetCurrentShareGroupResponseMapper(shareGroup, user.ID)); err != nil {
				log.Println(err)
				return err
			}
		}

		time.Sleep(time.Second * 1)
	}
}
