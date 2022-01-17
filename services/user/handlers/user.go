package handlers

import (
	"github.com/royleochan/recall/services/user/schemas"
	"github.com/royleochan/recall/services/user/userpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type Server struct {
	*mongo.Collection
	userpb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId, _ := primitive.ObjectIDFromHex(req.UserId)

	var user = schemas.User{}
	err := s.Collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{UserId: user.Id.Hex()}, nil
}
