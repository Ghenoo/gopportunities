package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	Router := gin.Default()

	// Initialize Routes
	InitializeRoutes(Router)

	// Run the server
	Router.Run("address:8080")
}
