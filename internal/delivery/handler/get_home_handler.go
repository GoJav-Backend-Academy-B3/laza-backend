package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/contract"
	"github.com/phincon-backend/laza/helper"
)

type HomeHandler struct {
}

func NewHomeHandler() contract.MainHandlerInterface {
	return &HomeHandler{}
}

// GetHandler implements contract.MainHandlerInterface.
func (h *HomeHandler) GetHandler() (method string, path string, handlerFuncs gin.HandlerFunc) {
	return http.MethodGet, "/", h.Handle
}

// Handle implements contract.MainHandlerInterface.
func (h *HomeHandler) Handle(c *gin.Context) {
	response := map[string]string{
		"repo": "",
		"demo": "",
		"docs": "",
	}

	helper.GetResponse(response, 200, true).Send(c)
}
