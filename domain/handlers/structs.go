package handlers

import (
	"github.com/gin-gonic/gin"
)

type HandlerStruct struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func (h HandlerStruct) GinHandlerFunc() (method string, path string, handlerFunc gin.HandlerFunc) {
	return h.Method, h.Path, h.HandlerFunc
}

func (h HandlerStruct) GinMiddlewares() []gin.HandlerFunc {
	return h.Middlewares
}
