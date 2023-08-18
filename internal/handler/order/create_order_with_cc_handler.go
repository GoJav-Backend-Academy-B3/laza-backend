package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

// Create order with cedit card godoc
// @Summary Create order with cedit card
// @Description Create order with cedit card payment
// @Tags order
// @Accept json
// @Produce json
// @Param wishlist body requests.OrderWithCC true "order request"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.OrderCreateResponse}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /order/cc [POST]
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

	orderResponse := response.Order{}
	orderResponse.FillFromEntity(order)

	result := response.OrderCreateResponse{
		Order:         orderResponse,
		PaymentMethod: *paymentMethod,
	}

	response := helper.GetResponse(result, http.StatusCreated, false)
	response.Send(c)
}
