package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/royleochan/recall/backend/controllers"
)

func InitFlashcardRoutes(router *gin.Engine) {

	flashcardRoutes := router.Group("/flashcard")
	{
		flashcardRoutes.GET("/:flashcard_id", controllers.GetFlashcardById)
	}
}
