package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// VerificationCode godoc
// @Summary Verification Email For User
// @Description Verification email for user
// @Tags auth
// @Accept json
// @Produce json
// @Param auth body requests.VerificationCode true "Verification Code"
// @Success 202 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/verification-code [post]
func (h *authHandler) verificationCode(c *gin.Context) {
	var request requests.VerificationCode
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.verificationCodeUser.Execute(request).Send(c)
}
