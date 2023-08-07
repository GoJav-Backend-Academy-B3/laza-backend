package bank

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

func (h *bankHandler) insert(c *gin.Context) {
	var request model.Bank
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.insertBank.Execute(request).Send(c)
}
