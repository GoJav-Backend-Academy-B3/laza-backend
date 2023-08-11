package auth

import (
	"github.com/gin-gonic/gin"
)

// VerificationEmail godoc
// @Summary Verification Email For User
// @Description Verification email for user
// @Tags auth
// @Accept json
// @Produce json
// @Param token query string true "Query Token"
// @Param email query string true "Query Email "
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/confirm [get]
func (h *authHandler) confirmEmail(c *gin.Context) {
	email := c.Query("email")
	token := c.Query("token")

	h.verifyEmailUser.Execute(email, token).Send(c)
}
