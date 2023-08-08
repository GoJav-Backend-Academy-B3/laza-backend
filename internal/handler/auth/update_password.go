package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// UpdatePassword godoc
// @Summary Update Password For User
// @Description Update password for user
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "Query Code"
// @Param email query string true "Query Email "
// @Param auth body requests.UpdatePassword true "Update Password"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/update-password/ [post]
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

	var request requests.UpdatePassword
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
