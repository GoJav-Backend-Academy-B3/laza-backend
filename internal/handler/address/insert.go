package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

func (h *addressHandler) PostAddressHandler(ctx *gin.Context) {
	var request requests.AddressRequest

	userAuth := ctx.MustGet("authID").(jwt.MapClaims)
	userId := uint64(userAuth["UserId"].(float64))

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	address, err := h.insert.AddAddress(request, userId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(address, http.StatusCreated, false).Send(ctx)

}
