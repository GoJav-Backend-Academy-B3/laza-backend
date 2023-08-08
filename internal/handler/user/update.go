package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// UpdateUser godoc
// @Summary Update User
// @Description Update for user
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param user formData requests.User true "user"
// @Param image formData file false "user"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=response.User}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /user/update [put]
func (h *userHandler) update(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 2*1024*1024)

	userId := c.MustGet("userId").(uint64)

	var request requests.UpdateUser
	if err := c.Bind(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.updateUser.Execute(userId, request).Send(c)
}
