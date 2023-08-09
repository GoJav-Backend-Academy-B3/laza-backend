package bank

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// GetByIdBank godoc
// @Summary Get bank identified by the given id
// @Description get the bank corresponding to the input Id
// @Tags bank
// @Accept json
// @Produce json
// @Param id path int true "ID of the bank"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Bank}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /bank/{id} [get]
func (h *bankHandler) getById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	idParse, _ := strconv.ParseUint(id, 10, 64)
	h.getByIdBank.Execute(idParse).Send(c)
}
