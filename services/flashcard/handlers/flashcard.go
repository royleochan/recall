package handlers

import (
	"context"

	"github.com/royleochan/recall/services/flashcard/flashcardpb"
	"github.com/royleochan/recall/services/flashcard/models"
	"gorm.io/gorm"
)

type Server struct {
	Db                     *gorm.DB
	FlashcardServiceServer flashcardpb.UnimplementedFlashcardServiceServer
}

type FlashcardResponse struct {
	models.Flashcard
	models.Board
}

func (s *Server) GetFlashcardById(ctx context.Context, req *flashcardpb.GetFlashcardByIdRequest) (*flashcardpb.GetFlashcardByIdResponse, error) {
	var res FlashcardResponse

	query := s.Db.Model(&models.Flashcard{}).Select("flashcards.id, flashcards.title, flashcards.answer, boards.name, boards.creator").Joins("join boards on boards.id = flashcards.board_id").Where("flashcards.id = ?", req.FlashcardId).First(&res)
	if query.Error != nil {
		return nil, query.Error
	}

	return &flashcardpb.GetFlashcardByIdResponse{FlashcardId: int64(res.Flashcard.ID), Title: res.Title, Answer: res.Answer, Board: res.Name, Creator: res.Creator}, nil
}

func (s *Server) GetFlashcards(req *flashcardpb.GetFlashcardsRequest, stream flashcardpb.FlashcardService_GetFlashcardsServer) error {
	return nil
}

func (s *Server) CreateFlashcard(ctx context.Context, req *flashcardpb.CreateFlashcardRequest) (*flashcardpb.Default, error) {
	return nil, nil
}

func (s *Server) UpdateFlashcard(ctx context.Context, req *flashcardpb.UpdateFlashcardRequest) (*flashcardpb.Default, error) {
	return nil, nil
}

func (s *Server) DeleteFlashcard(ctx context.Context, req *flashcardpb.DeleteFlashcardRequest) (*flashcardpb.Default, error) {
	return nil, nil
}
