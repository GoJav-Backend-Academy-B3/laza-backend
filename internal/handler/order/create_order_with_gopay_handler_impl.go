package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *orderHandler) CreateOrderWithGopay(c *gin.Context) {
	var orderRequest request.OrderWithGopay

	err := c.ShouldBind(&orderRequest)
	if err != nil {
		response := helper.GetResponse(err, http.StatusBadRequest, true)
		response.Send(c)
		return
	}

	userId := c.MustGet("userId").(uint64)

	order, gopay, err := h.createOrderWithGopayUsecase.Execute(userId, orderRequest.AddressId, orderRequest.CallbackUrl, orderRequest.Products)
	if err != nil {
		response := helper.GetResponse(err, http.StatusInternalServerError, true)
		response.Send(c)
		return
	}

	result := make(map[string]any)

	orderResponse := &response.GopayOrderResponse{}
	orderResponse.FillFromEntity(order)

	gopayPaymentResponse := &response.GopayPaymentResponse{}
	gopayPaymentResponse.FillFromEntity(gopay)
	result["order"] = orderResponse
	result["gopay"] = gopayPaymentResponse

	response := helper.GetResponse(result, http.StatusOK, false)
	response.Send(c)
}
