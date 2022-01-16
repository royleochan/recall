package handlers

import (
	"log"

	"github.com/royleochan/recall/services/user/userpb"
	"golang.org/x/net/context"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Println("Called GetUser Operation")

	return &userpb.GetUserResponse{UserId: req.UserId}, nil
}
