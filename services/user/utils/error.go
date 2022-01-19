package utils

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleMongoError(err error, entity string) error {
	if err == mongo.ErrNoDocuments {
		return status.Errorf(codes.NotFound, "No %s found", entity)
	}
	return err
}
