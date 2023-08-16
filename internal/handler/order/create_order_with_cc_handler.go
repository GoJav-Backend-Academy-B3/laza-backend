package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) CreateOrderWithCC(c *gin.Context) {
	var orderCCRequest requests.OrderWithCC

	err := c.ShouldBind(&orderCCRequest)
	if err != nil {
		response := helper.GetResponse(err.Error(), http.StatusBadRequest, true)
		response.Send(c)
		return
	}

	userId := c.MustGet("userId").(uint64)

	order, paymentMethod, err := h.createOrderWithCCUsecase.Execute(
		userId,
		orderCCRequest.AddressId,
		model.CreditCard{
			Id:           uint64(orderCCRequest.CreditCard.Id),
			CardNumber:   orderCCRequest.CreditCard.CardNumber,
			ExpiredMonth: orderCCRequest.CreditCard.ExpMonth,
			ExpiredYear:  orderCCRequest.CreditCard.ExpYear,
		},
		orderCCRequest.CreditCard.CVV,
	)
	if err != nil {
		response := helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		response.Send(c)
		return
	}

	result := make(map[string]any)

	//orderResponse := &response.CreditCardOrderResponse{}
	//orderResponse.FillFromEntity(order)

	result["order"] = order
	result["payment_method"] = paymentMethod

	response := helper.GetResponse(result, http.StatusCreated, false)
	response.Send(c)
}
