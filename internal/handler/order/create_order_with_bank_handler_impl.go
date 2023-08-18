package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

// Create order with bank godoc
// @Summary Create order with bank
// @Description Create order with bank payment
// @Tags order
// @Accept json
// @Produce json
// @Param wishlist body requests.OrderWithBank true "order request"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.OrderCreateResponse}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /order/bank [POST]
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

	orderResponse := response.Order{}
	orderResponse.FillFromEntity(order)

	result := response.OrderCreateResponse{
		Order:         orderResponse,
		PaymentMethod: *paymentMethod,
	}

	response := helper.GetResponse(result, http.StatusCreated, false)
	response.Send(c)
}
