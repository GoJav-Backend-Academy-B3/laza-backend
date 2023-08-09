package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) CreateOrderWithBank(c *gin.Context) {
	var orderRequest requests.OrderWithBank

	err := c.ShouldBind(&orderRequest)
	if err != nil {
		response := helper.GetResponse(err, http.StatusBadRequest, true)
		response.Send(c)
		return
	}

	userId := c.MustGet("userId").(uint64)

	order, bankDetails, err := h.createOrderWithBankUsecase.Execute(userId, orderRequest.AddressId, orderRequest.Bank, orderRequest.Products)
	if err != nil {
		response := helper.GetResponse(err, http.StatusInternalServerError, true)
		response.Send(c)
		return
	}

	result := make(map[string]any)

	orderResponse := &response.TransactionBankOrderResponse{}
	orderResponse.FillFromEntity(order)

	result["order"] = orderResponse
	result["va_detail"] = bankDetails

	response := helper.GetResponse(result, http.StatusCreated, false)
	response.Send(c)
}
