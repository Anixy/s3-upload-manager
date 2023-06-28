package routers

import (
	"github.com/Anixy/s3-upload-manager/handler"

	"github.com/gin-gonic/gin"
)

func Set(router *gin.Engine, handler handler.Handler) (err error) {
	router.GET("/ping", handler.Ping())
	return
}