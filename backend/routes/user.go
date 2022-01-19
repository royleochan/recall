package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/royleochan/recall/backend/controllers"
)

func InitUserRoutes(router *gin.Engine) {

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/:user_id", controllers.GetUserById)
	}
}
