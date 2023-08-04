package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/cmd/server/provider"
	"github.com/phincon-backend/laza/domain/handlers"
)

func NewServerGin() *gin.Engine {
	r := gin.Default()

	var server []handlers.HandlerInterface
	server = append(server,
		provider.NewHomeHandler(),
		provider.NewProductsHandler(),
	)

	for _, v := range server {
		handlers := v.GetHandlers()
		for _, handler := range handlers {
			method := handler.Method
			path := handler.Path
			hfunc := handler.HandlerFunc
			r.Handle(method, path, hfunc...)
		}
	}

	return r
}
