package order

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (h *orderHandler) GetOrderByUser(c *gin.Context) {

	// Get userId
	userId := c.MustGet("userId").(uint64)

	orders, err := h.getAllByUser.Execute(userId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	result := make([]response.Order, 0)
	for _, order := range orders {
		temp := response.Order{}
		temp.FillFromEntity(&order)
		result = append(result, temp)
	}

	helper.GetResponse(result, 200, false).Send(c)
	return
}
