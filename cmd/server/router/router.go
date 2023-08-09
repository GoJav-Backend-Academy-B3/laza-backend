package router

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/cmd/server/provider"
	_ "github.com/phincon-backend/laza/docs"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r := gin.Default()
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

	auth := r.Group("").Use(middleware.AuthMiddleware())
	adminAuth := r.Group("").Use(middleware.AuthMiddleware(), middleware.AdminRoleMiddleware())

	for _, v := range server {
		handlers := v.GetHandlers()
		for _, handler := range handlers {
			_, path, _ := handler.GinHandlerFunc()
			if noAuth(path) {
				r.Handle(handler.GinHandlerFunc())
			} else if strings.Contains(path, "/auth") {
				r.Handle(handler.GinHandlerFunc())
			} else if roleAdmin(path) {
				adminAuth.Handle(handler.GinHandlerFunc())
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
	noAuthList = append(noAuthList, "/login")
	noAuthList = append(noAuthList, "/register")
	noAuthList = append(noAuthList, "/login-google")
	noAuthList = append(noAuthList, "/login-google/callback")
	noAuthList = append(noAuthList, "/products")
	noAuthList = append(noAuthList, "/products/:id")
	noAuthList = append(noAuthList, "/size")
	noAuthList = append(noAuthList, "/size/:id")
	noAuthList = append(noAuthList, "/category")
	noAuthList = append(noAuthList, "/brand")
	noAuthList = append(noAuthList, "/brand/:id")
	noAuthList = append(noAuthList, "/brand/search")
	for _, item := range noAuthList {
		if strings.EqualFold(item, url) {
			return true
		}
	}
	return false
}

var roleAdminList = make([]string, 0)
func roleAdmin(url string) bool {
	roleAdminList = append(roleAdminList, "/user")
	roleAdminList = append(roleAdminList, "/user/")
	for _, item := range roleAdminList {
		if strings.EqualFold(item, url) {
			return true
		}
	}
	return false
}
