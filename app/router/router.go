package router

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/app/provider"
	"github.com/phincon-backend/laza/domain/contract"
)

func NewServerGin() *gin.Engine {
	r := gin.Default()

	var server []contract.MainHandlerInterface
	server = append(server, provider.NewHomeHandler())

	for _, v := range server {
		r.Handle(v.GetHandler())
	}

	return r
}
