package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

func (h *userHandler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.updateUser.Excute(uint64(c.GetInt64(id)), request).Send(c)
}
