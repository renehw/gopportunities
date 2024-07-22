package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST /api/v1/opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET /api/v1/opening",
	})
}
func DeleateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE /api/v1/opening",
	})
}
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT /api/v1/opening",
	})
}
func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET /api/v1/openings",
	})
}
