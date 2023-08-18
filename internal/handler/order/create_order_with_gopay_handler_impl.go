package order

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// Create order with gopay godoc
// @Summary Create order with gopay
// @Description Create order with gopay payment
// @Tags order
// @Accept json
// @Produce json
// @Param wishlist body requests.OrderWithGopay true "order request"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.OrderCreateResponse}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /order/gopay [POST]
func (h *orderHandler) CreateOrderWithGopay(c *gin.Context) {
	var orderRequest requests.OrderWithGopay

	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	userId := c.MustGet("userId").(uint64)

	order, paymentMethod, err := h.createOrderWithGopayUsecase.Execute(userId, orderRequest.AddressId, orderRequest.CallbackUrl)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	orderResponse := response.Order{}
	orderResponse.FillFromEntity(order)

	result := response.OrderCreateResponse{
		Order:         orderResponse,
		PaymentMethod: *paymentMethod,
	}

	helper.GetResponse(result, http.StatusCreated, false).Send(c)
}
