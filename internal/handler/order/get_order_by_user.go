package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) GetOrderById(c *gin.Context) {

	// Get limit and offset query string
	orderId := c.Param("order_id")

	if orderId == "" {
		helper.GetResponse("missing order id", http.StatusInternalServerError, true).Send(c)
		return
	}

	order, err := h.getById.Execute(orderId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(order, 200, false).Send(c)
	return

}
