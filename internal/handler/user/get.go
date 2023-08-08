package user

import "github.com/gin-gonic/gin"

// GetAllUser godoc
// @Summary Get All User
// @Description Get all user
// @Tags user
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=response.User}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /user [get]
func (h *userHandler) get(c *gin.Context) {
	h.getAllUser.Execute().Send(c)
}
