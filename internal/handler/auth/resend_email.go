package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// ResendEmail godoc
// @Summary Resend Email For Verify Email User
// @Description Resend Email for verify email user
// @Tags auth
// @Accept json
// @Produce json
// @Param auth body requests.Email true "Resend Email"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/resend-verify [post]
func (h *authHandler) resendEmail(c *gin.Context) {
	var request requests.Email
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
