package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

func (h *addressHandler) UpdateAddressHandler(ctx *gin.Context) {
	var request requests.AddressRequest

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	userAuth := ctx.MustGet("authID").(jwt.MapClaims)
	userId := uint64(userAuth["UserId"].(float64))

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		if err != nil {
			helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
			return
		}
	}

	err = h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	addresses, err := h.update.UpdateAddressById(id, userId, request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(addresses, http.StatusOK, false).Send(ctx)

}
