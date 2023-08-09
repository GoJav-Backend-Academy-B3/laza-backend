package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Register godoc
// @Summary Register
// @Description Register for user
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param user formData requests.Register true "user"
// @Param image formData file false "user"
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=response.User}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /register [post]
func (h *authHandler) register(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 2*1024*1024)

	var request requests.Register
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
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
