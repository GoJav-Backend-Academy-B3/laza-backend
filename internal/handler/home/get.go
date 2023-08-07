package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *HomeHandler) get(c *gin.Context) {
	response_map := map[string]string{
		"repo": "",
		"demo": "",
		"docs": "",
	}
	helper.GetResponse(response_map, 200, false).Send(c)
}
