package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) login(c *gin.Context) {
	var request request.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.loginUser.Execute(request).Send(c)
}
