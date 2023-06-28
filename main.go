package main

import (
	"github.com/Anixy/s3-upload-manager/handler"
	"github.com/Anixy/s3-upload-manager/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := handler.NewHandler()
	router := gin.Default()
	routers.Set(router, handler)
	router.Run(":8080")
}
