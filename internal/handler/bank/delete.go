package bank

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *bankHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	idParse, _ := strconv.ParseUint(id, 10, 64)
	h.deleteBank.Execute(idParse).Send(c)
}
