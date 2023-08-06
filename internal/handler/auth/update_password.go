package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) updatePassword(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")

	if email == "" && code == "" {
		helper.GetResponse("email and code are both empty", 400, true).Send(c)
		return
	} else if email == "" {
		helper.GetResponse("email empty", 400, true).Send(c)
		return
	} else if code == "" {
		helper.GetResponse("code empty", 400, true).Send(c)
		return
	}

	var request request.UpdatePassword
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.updatePasswordUser.Execute(email, code, request).Send(c)
}
