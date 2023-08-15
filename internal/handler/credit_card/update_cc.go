package credit_card

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Update Credit Card godoc
// @Summary Update Credit Card
// @Description Credit Card
// @Tags creditcard
// @Accept json
// @Produce json
// @Param id path int true "ID of the credit card"
// @Param creditcard body requests.CreditCardRequest true "update credit card"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.CreditCardResponse}
// @Failure 404 {object} helper.Response{status=string,description=string,isError=bool}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /credit-card/{id} [PUT]
func (h *getCreditCardHandler) Update(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	ccId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}

	requestBody := requests.CreditCardRequest{}
	if err := c.Bind(&requestBody); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true)
		return
	}

	_result, statusCode, err := h.updateCcUc.Execute(userId, ccId, requestBody)
	if err != nil {
		helper.GetResponse(err.Error(), statusCode, true).Send(c)
		return
	}

	helper.GetResponse(_result, 200, false).Send(c)

}
