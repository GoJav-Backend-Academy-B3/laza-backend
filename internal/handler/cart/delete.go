package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Delete Cart
// @Summary Delete Cart
// @Description Delete cart with product id
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=string}
// @Failure 401 {object} helper.Response{status=string,description=string,isError=bool}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /products/{id}/carts [delete]
func (ct *CartHandler) Delete(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	var requestBody requests.CartRequest
	err := ctx.Bind(&requestBody)

	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(ctx)
		return
	}

	_result, err := ct.deleteCartUc.Execute(userId, requestBody)
	if ex, ok := err.(validator.ValidationErrors); ok {
		helper.GetResponse(ex.Error(), 400, true).Send(ctx)
		return
	}

	helper.GetResponse(_result, 200, false).Send(ctx)

}
