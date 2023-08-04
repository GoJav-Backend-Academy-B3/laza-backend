package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *userHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	h.deleteUser.Excute(uint64(c.GetInt64(id))).Send(c)
}
