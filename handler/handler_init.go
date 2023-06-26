package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Ping() gin.HandlerFunc
}


type handlerImpl struct {
	
}

func NewHandler() Handler {
	return &handlerImpl{}
}