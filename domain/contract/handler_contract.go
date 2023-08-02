package contract

import (
	"github.com/gin-gonic/gin"
)

type MainHandlerInterface interface {
	Handle(c *gin.Context)
	GetHandler() (method string, path string, handlerFuncs gin.HandlerFunc)
}
