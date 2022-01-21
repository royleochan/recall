package handlers

import (
	"context"

	"github.com/royleochan/recall/services/flashcard/database"
	"github.com/royleochan/recall/services/flashcard/flashcardpb"
	"github.com/royleochan/recall/services/flashcard/models"
)

var db = database.DBCon

func GetFlashcardById(ctx context.Context, req *flashcardpb.GetFlashcardByIdRequest) (*flashcardpb.GetFlashcardByIdResponse, error) {
	var flashcard models.Flashcard
	result := db.First(&flashcard, req.FlashcardId)

	if result.Error != nil {
		return nil, result.Error
	}

	println(result)

	return &flashcardpb.GetFlashcardByIdResponse{}, nil
}
