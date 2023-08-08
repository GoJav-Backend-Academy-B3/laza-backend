package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) CreateOrderWithBank(c *gin.Context) {
	var orderRequest request.OrderWithBank

	err := c.ShouldBind(&orderRequest)
	if err != nil {
		response := helper.GetResponse(err, http.StatusBadRequest, true)
		response.Send(c)
	}

	userId := c.MustGet("userId").(uint64)

	order, bankDetails, err := h.createOrderWithBankUsecase.Execute(userId, orderRequest.AddressId, orderRequest.Bank, orderRequest.Products)
	if err != nil {
		response := helper.GetResponse(err, http.StatusBadRequest, true)
		response.Send(c)
	}

	result := make(map[string]any)

	result["order"] = order
	result["bankDetails"] = bankDetails

	response := helper.GetResponse(result, http.StatusOK, false)
	response.Send(c)
}
