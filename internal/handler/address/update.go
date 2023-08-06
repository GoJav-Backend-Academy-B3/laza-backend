package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

func (h *addressHandler) UpdateAddressHandler(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	var request requests.AddressRequest

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

	addresses, err := h.update.UpdateAddressById(id, request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(addresses, http.StatusOK, false).Send(ctx)

}
