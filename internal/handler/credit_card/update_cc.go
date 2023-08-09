package credit_card

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

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

	h.updateCcUc.Execute(userId, ccId, requestBody).Send(c)

}
