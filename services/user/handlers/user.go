package handlers

import (
	"github.com/royleochan/recall/services/user/schemas"
	"github.com/royleochan/recall/services/user/userpb"
	"github.com/royleochan/recall/services/user/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

const ENTITY = "user"

type Server struct {
	Collection        *mongo.Collection
	UserServiceServer userpb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId, _ := primitive.ObjectIDFromHex(req.UserId)

	var user = schemas.User{}
	err := s.Collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)

	if err != nil {
		return nil, utils.HandleMongoError(err, ENTITY)
	}

	return &userpb.GetUserResponse{UserId: user.Id.Hex()}, nil
}
