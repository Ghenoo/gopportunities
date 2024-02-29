package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	Router := gin.Default()

	// Initialize Routes
	initializeRoutes(Router)

	// Run the server
	Router.Run("address:8080")
}
