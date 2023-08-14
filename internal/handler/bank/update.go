package bank

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// UpdateBank godoc
// @Summary put Details
// @Description put details of address
// @Tags bank
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "ID of the bank"
// @Param bank formData requests.BankRequest true "update bank"
// @Param image formData file true "bank"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Bank}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /bank/{id} [put]
func (h *bankHandler) update(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	var request requests.BankRequest
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	idParse, _ := strconv.ParseUint(id, 10, 64)

	h.updateBank.Execute(idParse, request).Send(c)
}
