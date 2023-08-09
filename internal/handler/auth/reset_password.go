package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// ResetPassword godoc
// @Summary Reset Password For User
// @Description Reset password for user
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "Query Code"
// @Param email query string true "Query Email "
// @Param auth body requests.ResetPassword true "Reset Password"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/reset-password/ [post]
func (h *authHandler) resetPassword(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")

	var request requests.ResetPassword
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.resetPasswordUser.Execute(email, code, request).Send(c)
}
