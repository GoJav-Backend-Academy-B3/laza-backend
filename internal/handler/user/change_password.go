package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// ChangePasswordUser godoc
// @Summary Change Password User
// @Description Change passsword user
// @Tags user
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=map[string]string}
// @Failure 400 {object} helper.Response{status=string,description=string,isError=bool}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /user/change-password [put]
func (h *userHandler) changePassword(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)

	var request requests.ChangePassword
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}
	h.changePasswordUser.Execute(userId, request).Send(c)
}
