package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) register(c *gin.Context) {
	var request requests.User
	if err := c.Bind(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	url, err := helper.UploadImage(c)
	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(c)
		return
	}

	request.Image = url
	h.registerUser.Execute(request).Send(c)
}
