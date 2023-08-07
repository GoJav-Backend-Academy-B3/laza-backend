package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) verifyEmail(c *gin.Context) {
	email := c.Query("email")
	token := c.Query("token")

	if email == "" && token == "" {
		helper.GetResponse("email and token are both empty", 400, true).Send(c)
		return
	} else if email == "" {
		helper.GetResponse("email empty", 400, true).Send(c)
		return
	} else if token == "" {
		helper.GetResponse("token empty", 400, true).Send(c)
		return
	}

	h.verifyEmailUser.Execute(email, token).Send(c)
}
