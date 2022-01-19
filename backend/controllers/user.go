package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/royleochan/recall/backend/datasources"
	"github.com/royleochan/recall/backend/utils"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

func GetUserById(c *gin.Context) {
	user, err := datasources.GetUserById(c.Param("user_id"))

	if err != nil {
		utils.HandleGrpcError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &User{Id: user.UserId, Username: user.Username, Email: user.Email, Name: user.Name, Avatar: user.Avatar}})
	}
}
