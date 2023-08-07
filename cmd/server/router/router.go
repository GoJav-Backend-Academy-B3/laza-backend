package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/cmd/server/provider"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/middleware"
	"strings"
)

func NewServerGin() *gin.Engine {
	r := gin.Default()

	var server []handlers.HandlerInterface
	server = append(server,
		provider.NewHomeHandler(),
		provider.NewAuthHandler(),
		provider.NewUserHandler(),
		provider.NewProductsHandler(),
		provider.NewWishListsHandler(),
		provider.NewCartHandler(),
		provider.NewViewProductByBrandHandler(),
		provider.NewFacebookAuthHandler(),
	)
	auth := r.Group("").Use(middleware.AuthMiddleware())
	for _, v := range server {
		handlers := v.GetHandlers()
		for _, handler := range handlers {
			_, path, _ := handler.GinHandlerFunc()
			if noAuth(path) {
				r.Handle(handler.GinHandlerFunc())
			} else if strings.Contains(path, "/auth") {
				r.Handle(handler.GinHandlerFunc())
			} else {
				auth.Handle(handler.GinHandlerFunc())
			}
		}
	}

	return r
}

var noAuthList = make([]string, 0)

func noAuth(url string) bool {
	noAuthList = append(noAuthList, "/")
	for _, item := range noAuthList {
		if strings.EqualFold(item, url) {
			return true
		}
	}
	return false
}
