package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/royleochan/recall/backend/datasources"
	"github.com/royleochan/recall/backend/utils"
)

type Board struct {
	Name    string `json:"name"`
	Creator string `json:"creator"`
}

type Flashcard struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Answer string `json:"answer"`
	Board  Board  `json:"board"`
}

func GetFlashcardById(c *gin.Context) {
	flashcardId, _ := strconv.Atoi(c.Param("flashcard_id"))
	flashcard, err := datasources.GetFlashcardById(int64(flashcardId))

	if err != nil {
		utils.HandleGrpcError(c, err)
	} else {
		// board := &Board{Name: flashcard.Board.Name, Creator: flashcard.Board.Creator}
		c.JSON(http.StatusOK, gin.H{"data": &Flashcard{Id: flashcard.FlashcardId, Title: flashcard.Title, Answer: flashcard.Answer}})
	}
}
