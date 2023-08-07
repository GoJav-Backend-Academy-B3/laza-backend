package user

import (
	"github.com/gin-gonic/gin"
)

func (h *userHandler) getById(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	h.getByIdUser.Execute(userId).Send(c)
}
