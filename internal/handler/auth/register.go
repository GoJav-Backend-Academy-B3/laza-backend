package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Register godoc
// @Summary Register
// @Description Register for user
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param auth body requests.Register true "Register"
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=response.User}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /register [post]
func (h *authHandler) register(c *gin.Context) {
	var request requests.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.registerUser.Execute(request).Send(c)
}
