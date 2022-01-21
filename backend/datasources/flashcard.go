package datasources

import (
	"context"
	"errors"

	"github.com/royleochan/recall/services/flashcard/flashcardpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var flashcardGrpcService flashcardpb.FlashcardServiceClient

func InitFlashcardServiceConn() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("connection to flashcard gRPC service failed")
	}

	flashcardGrpcService = flashcardpb.NewFlashcardServiceClient(conn)

	return conn, nil
}

func GetFlashcardById(id int64) (*flashcardpb.GetFlashcardByIdResponse, error) {
	response, err := flashcardGrpcService.GetFlashcardById(context.Background(), &flashcardpb.GetFlashcardByIdRequest{FlashcardId: id})
	if err != nil {
		return nil, err
	}

	return response, nil
}
