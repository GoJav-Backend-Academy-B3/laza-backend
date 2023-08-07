package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) resendEmail(c *gin.Context) {
	var request request.Resend
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.resendEmailUser.Execute(request.Email).Send(c)
}
