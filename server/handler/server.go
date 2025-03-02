package handler

import pb "main/go_protocol_buffer"

type Server struct {
	pb.UnimplementedServiceServer
}
