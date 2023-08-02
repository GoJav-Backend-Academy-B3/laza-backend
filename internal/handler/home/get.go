package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/response"
)

func (h *HomeHandler) get(c *gin.Context) {
	response_map := map[string]string{
		"repo": "",
		"demo": "",
		"docs": "",
	}
	response.GetResponse(response_map, http.StatusOK, false)
}
