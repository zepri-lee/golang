package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthhenticated",
		})
		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token is not valid",
		})
		return
	}

	ctx.Next()
}
