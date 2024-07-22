package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Inicializa o Servidor na porta 8080
	router.Run(":8080")
}