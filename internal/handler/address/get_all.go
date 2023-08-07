package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/phincon-backend/laza/helper"
)

func (h *addressHandler) GetAllAddressByUserIdHandler(ctx *gin.Context) {
	userAuth := ctx.MustGet("authID").(jwt.MapClaims)
	userId := uint64(userAuth["UserId"].(float64))

	addresses, err := h.get.GetAllAddressByUserId(userId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(addresses, http.StatusOK, false).Send(ctx)

}
