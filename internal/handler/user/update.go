package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

func (h *userHandler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	var request request.User
	if err := c.Bind(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	url, err := helper.UploadImage(c)
	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(c)
		return
	}

	request.Image = url
	idParse, _ := strconv.ParseUint(id, 10, 64)
	h.updateUser.Execute(idParse, request).Send(c)
}
