package handlers

import (
	"github.com/gin-gonic/gin"
)

type HandlerStruct struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func (h HandlerStruct) HttpHandlerFunc() (method string, path string, handlerFunc gin.HandlerFunc) {
	return h.Method, h.Path, h.HandlerFunc
}
