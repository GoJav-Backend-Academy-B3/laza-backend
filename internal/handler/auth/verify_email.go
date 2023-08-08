package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// VerifyEmail godoc
// @Summary Verify Email For User
// @Description Verify email for user
// @Tags auth
// @Accept json
// @Produce json
// @Param token query string true "Query Token"
// @Param email query string true "Query Email "
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/verify-email/ [get]
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
