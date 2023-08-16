package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) CreateOrderWithBank(c *gin.Context) {
	var orderRequest requests.OrderWithBank

	err := c.ShouldBind(&orderRequest)
	if err != nil {
		response := helper.GetResponse(err.Error(), http.StatusBadRequest, true)
		response.Send(c)
		return
	}

	userId := c.MustGet("userId").(uint64)

	order, paymentMethod, err := h.createOrderWithBankUsecase.Execute(userId, orderRequest.AddressId, orderRequest.Bank)
	if err != nil {
		response := helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		response.Send(c)
		return
	}

	result := make(map[string]any)

	//orderResponse := &response.TransactionBankOrderResponse{}
	//orderResponse.FillFromEntity(order)

	result["order"] = order
	result["payment_method"] = paymentMethod

	response := helper.GetResponse(result, http.StatusCreated, false)
	response.Send(c)
}
