package user

import (
	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary Delete User
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /user/delete [delete]
func (h *userHandler) delete(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	h.deleteUser.Execute(userId).Send(c)
}
