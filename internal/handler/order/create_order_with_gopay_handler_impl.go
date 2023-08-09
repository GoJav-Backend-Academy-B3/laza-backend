package order

import (
	"fmt"
	"github.com/phincon-backend/laza/domain/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *orderHandler) CreateOrderWithGopay(c *gin.Context) {
	var orderRequest requests.OrderWithGopay

	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		fmt.Println("error binding: ", err)
		response := helper.GetResponse(err.Error(), http.StatusBadRequest, true)
		response.Send(c)
		return
	}

	userId := c.MustGet("userId").(uint64)

	order, gopay, err := h.createOrderWithGopayUsecase.Execute(userId, orderRequest.AddressId, orderRequest.CallbackUrl, orderRequest.Products)
	if err != nil {
		fmt.Println("error create order: ", err)
		response := helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
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

	response := helper.GetResponse(result, http.StatusCreated, false)
	response.Send(c)
}
