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
		provider.NewUserHandler(),
		provider.NewProductsHandler(),
		provider.NewWishListsHandler(),
	)

	for _, v := range server {
		handlers := v.GetHandlers()
		for _, handler := range handlers {
			r.Handle(handler.GinHandlerFunc())
		}
	}

	return r
}
