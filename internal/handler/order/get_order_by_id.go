package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

// Get order by id godoc
// @Summary Get order by id
// @Description Get detail order by id order
// @Tags order
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.OrderResponse}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /orders [GET]
func (h *orderHandler) GetOrderById(c *gin.Context) {

	// Get limit and offset query string
	orderId := c.Param("order_id")

	if orderId == "" {
		helper.GetResponse("missing order id", http.StatusInternalServerError, true).Send(c)
		return
	}

	order, productDetails, err := h.getById.Execute(orderId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	orderResponse := response.Order{}
	orderResponse.FillFromEntity(&order)

	result := response.OrderResponse{
		Order:         orderResponse,
		ProductDetail: productDetails,
	}

	helper.GetResponse(result, 200, false).Send(c)
	return
}
