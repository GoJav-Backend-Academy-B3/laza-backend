package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

func (h *userHandler) update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id <= 0 {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	var request request.User
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
	h.updateUser.Execute(id, request).Send(c)
}
