package handler

import (
	"context"
	"fmt"
	"main/database"
	pb "main/go_protocol_buffer"
	"main/server/crud"
	"main/server/view"
)

func (s *Server) GetShareGroupByLinkKey(ctx context.Context, req *pb.GetShareGroupByLinkKeyRequest) (
	*pb.GetShareGroupByLinkKeyResponse, error,
) {
	db := database.GetDB()

	shareGroup := crud.GetShareGroupByLinkKey(db, req.LinkKey)
	if shareGroup == nil {
		return nil, fmt.Errorf("error: the linkKey is invalid")
	}

	return view.GetShareGroupLinkKeyResponseMapper(shareGroup), nil
}
