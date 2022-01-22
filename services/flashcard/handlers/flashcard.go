package handlers

import (
	"context"

	"github.com/royleochan/recall/services/flashcard/flashcardpb"
	"gorm.io/gorm"
)

type Server struct {
	Db                     *gorm.DB
	FlashcardServiceServer flashcardpb.UnimplementedFlashcardServiceServer
}

func (s *Server) GetFlashcardById(ctx context.Context, req *flashcardpb.GetFlashcardByIdRequest) (*flashcardpb.GetFlashcardByIdResponse, error) {
	var flashcard flashcardpb.GetFlashcardByIdResponse
	result := s.Db.First(&flashcard, req.FlashcardId).Select("flashcards.id, flashcards.title, flashcards.answer, boards.name, boards.creator").Joins("join boards on boards.id = flashcards.board_id")
	if result.Error != nil {
		return nil, result.Error
	}

	return &flashcard, nil
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
