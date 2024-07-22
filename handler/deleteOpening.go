package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE /api/v1/opening",
	})
}
