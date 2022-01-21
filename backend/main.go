package main

import (
	"github.com/gin-gonic/gin"
	"github.com/royleochan/recall/backend/datasources"
	"github.com/royleochan/recall/backend/routes"
)

func main() {
	// Connect to datasources
	user_conn, _ := datasources.InitUserServiceConn()
	flashcard_conn, _ := datasources.InitFlashcardServiceConn()
	defer user_conn.Close()
	defer flashcard_conn.Close()

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	// Initialize custom routes
	router := gin.Default()
	routes.InitUserRoutes(router)
	routes.InitFlashcardRoutes(router)

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
