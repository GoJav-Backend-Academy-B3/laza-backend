package user

import (
	"github.com/gin-gonic/gin"
)

// ProfileUser godoc
// @Summary Profile User
// @Description Profile user
// @Tags user
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=response.User}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /user/profile [get]
func (h *userHandler) getById(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	h.getByIdUser.Execute(userId).Send(c)
}
