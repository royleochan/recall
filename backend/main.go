package main

import (
	"github.com/gin-gonic/gin"
	"github.com/royleochan/recall/backend/datasources"
	"github.com/royleochan/recall/backend/routes"
)

func main() {
	// Connect to datasources
	user_conn, _ := datasources.InitUserServiceConn()
	defer user_conn.Close()

	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	// Initialize custom routes
	router := gin.Default()
	routes.InitUserRoutes(router)

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
