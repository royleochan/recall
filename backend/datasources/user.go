package datasources

import (
	"context"
	"errors"

	"github.com/royleochan/recall/services/user/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userGrpcService userpb.UserServiceClient

func InitUserServiceConn() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("connection to user gRPC service failed")
	}

	userGrpcService = userpb.NewUserServiceClient(conn)

	return conn, nil
}

func GetUserById(id string) (*userpb.GetUserResponse, error) {
	response, err := userGrpcService.GetUser(context.Background(), &userpb.GetUserRequest{UserId: id})
	if err != nil {
		return nil, err
	}

	return response, nil
}
