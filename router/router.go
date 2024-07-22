package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initializa Router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server on :8080
	router.Run(":8080")
}
