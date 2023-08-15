package credit_card

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Add Credit Card godoc
// @Summary Add Credit Card
// @Description Credit Card
// @Tags creditcard
// @Accept json
// @Produce json
// @Param creditcard body requests.CreditCardRequest true "create credit card"
// @Security JWT
// @Success 201 {object} helper.Response{status=string,isError=bool,data=response.CreditCardResponse}
// @Failure 400 {object} helper.Response{status=string,description=string,isError=bool}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /credit-card [POST]
func (h *getCreditCardHandler) Add(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)

	var requestBody requests.CreditCardRequest
	if err := c.Bind(&requestBody); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	_result, statusCode, err := h.addCcUc.Execute(userId, requestBody)

	if err != nil {
		helper.GetResponse(err.Error(), statusCode, true).Send(c)
		return
	}

	helper.GetResponse(_result, 201, false).Send(c)
}
