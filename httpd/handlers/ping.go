package handlers

import "github.com/gin-gonic/gin"

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, map[string]bool{
			"success": true,
		})
	}
}
