package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *userHandler) getById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id <= 0 {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	h.getByIdUser.Execute(id).Send(c)
}
