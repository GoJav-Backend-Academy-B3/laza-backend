package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/cmd/server/provider"

	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Laza
// @version 1.0
// @description This is a Final Project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        X-Auth-Token
// @description	How to input in swagger : 'Bearer <insert_your_token_here>'
func NewServerGin() *gin.Engine {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var server []handlers.HandlerInterface
	server = append(server,
		provider.NewHomeHandler(),
		provider.NewAuthHandler(),
		provider.NewUserHandler(),
		provider.NewSizeHandler(),
		provider.NewProductsHandler(),
		provider.NewWishListsHandler(),
		provider.NewCartHandler(),
		provider.NewBankHandler(),
		provider.NewReviewHandler(),
		provider.NewViewProductByBrandHandler(),
		provider.NewFacebookAuthHandler(),
		provider.NewtwitterAuthHandler(),
		provider.NewAddressesHandler(),
		provider.NewOrderHandler(),
		provider.NewCategoryHandler(),
		provider.NewBrandHandler(),
		provider.NewcreditCardHandler(),
	)
	r.Use(middleware.LoggerMiddleware())

	for _, v := range server {
		handlersList := v.GetHandlers()
		for _, handler := range handlersList {
			method, path, handlerFunc := handler.GinHandlerFunc()
			r.Handle(method, path, append(handler.GinMiddlewares(), handlerFunc)...)
		}
	}

	return r
}
