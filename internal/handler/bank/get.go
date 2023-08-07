package bank

import "github.com/gin-gonic/gin"

func (h *bankHandler) get(c *gin.Context) {
	h.getAllBank.Execute().Send(c)
}
