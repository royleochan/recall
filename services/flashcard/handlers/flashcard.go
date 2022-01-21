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

func (s *Server) GetFlashcardById(ctx context.Context, req *flashcardpb.GetFlashcardByIdRequest) (*flashcardpb.GetFlashcardByIdResponse, error) {
	var flashcard models.Flashcard
	result := s.Db.First(&flashcard, req.FlashcardId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &flashcardpb.GetFlashcardByIdResponse{FlashcardId: int64(flashcard.ID), Title: flashcard.Title, Answer: flashcard.Answer, Board: nil}, nil
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
