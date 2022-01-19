package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcError(c *gin.Context, err error) {
	status, _ := status.FromError(err)
	errorCode := status.Code()
	errorMessage := status.Message()

	if errorCode == codes.NotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
	}
}
