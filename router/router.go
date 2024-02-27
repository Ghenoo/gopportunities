package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Incializa o Router (r) utilizando as configurações Default do gin
	Router := gin.Default()
	// Definindo a rota /ping
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Inicia o servidor na porta 8080
	Router.Run("address:8080")
}
