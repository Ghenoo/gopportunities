package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twitchyliquid64/golang-asm/obj"
)

func initializeRoutes(router *gin.Engine) {
	v1 *gin.RouterGroup := router.Group(relativePath: "/api/v1")
	{
		v1.GET(relativePath: "/opening", func(ctx *gin.Context) {	
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "GET Opening",
			 })
		 })
		 v1.POST(relativePath: "/opening", func(ctx *gin.Context) {	
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "POST Opening",
			 })
		 })
		 v1.DELETE(relativePath: "/opening", func(ctx *gin.Context) {	
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "DELETE Opening",
			 })
		 })
		 v1.PUT(relativePath: "/opening", func(ctx *gin.Context) {	
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "PUT Opening",
			 })
		 })
		 v1.GET(relativePath: "/openings", func(ctx *gin.Context) {	
			ctx.JSON(http.StatusOK,gin.H{
				"msg": "GET Opening",
			 })
		 })
	}
}
