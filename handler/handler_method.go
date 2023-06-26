package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h * handlerImpl) Ping() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}