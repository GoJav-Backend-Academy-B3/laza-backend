package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// ForgetPassword godoc
// @Summary Forget Password For User
// @Description Forget password for user
// @Tags auth
// @Accept json
// @Produce json
// @Param auth body requests.Email true "Forget Password"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/forget-password [post]
func (h *authHandler) forgotPassword(c *gin.Context) {
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

	h.forgotPasswordUser.Execute(request.Email).Send(c)
}
