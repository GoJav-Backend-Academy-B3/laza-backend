package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Update Cart godoc
// @Summary Update Cart
// @Description Update cart
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=model.Cart}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /products/{id}/carts [PUT]
func (h *CartHandler) Update(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	var requestBody requests.CartRequest
	err := ctx.Bind(&requestBody)

	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(ctx)
		return
	}

	_result, err := h.updateCartUc.Execute(userId, requestBody)
	if err, ok := err.(validator.ValidationErrors); ok {
		helper.GetResponse(err.Error(), 400, true).Send(ctx)
		return
	}

	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(ctx)
		return
	}

	helper.GetResponse(_result, 200, false).Send(ctx)
}
