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
	res, err := datasources.GetFlashcardById(int64(flashcardId))

	if err != nil {
		utils.HandleGrpcError(c, err)
	} else {
		board := &Board{Name: res.Board, Creator: res.Creator}
		c.JSON(http.StatusOK, gin.H{"data": &Flashcard{Id: res.FlashcardId, Title: res.Title, Answer: res.Answer, Board: *board}})
	}
}
