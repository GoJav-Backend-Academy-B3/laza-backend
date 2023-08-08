package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) CreateOrderWithGopay(c *gin.Context) {
	var orderRequest request.OrderWithGopay

	err := c.ShouldBind(&orderRequest)
	if err != nil {
		response := helper.GetResponse(err, http.StatusBadRequest, true)
		response.Send(c)
	}

	userId := c.MustGet("userId").(uint64)

	order, gopay, err := h.createOrderWithGopayUsecase.Execute(userId, orderRequest.AddressId, orderRequest.CallbackUrl, orderRequest.Products)
	if err != nil {
		return
	}

	result := make(map[string]any)

	result["order"] = order
	result["gopay"] = gopay

	response := helper.GetResponse(result, http.StatusOK, false)
	response.Send(c)
}
