package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// UpdateAddress godoc
// @Summary put Details
// @Description put details of address
// @Tags address
// @Accept json
// @Produce json
// @Param address body requests.AddressRequest true "update address"
// @Param id path int true "ID of the address"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Address}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Error 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /address/{id} [put]
func (h *addressHandler) UpdateAddressHandler(ctx *gin.Context) {
	var request requests.AddressRequest

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	userId := ctx.MustGet("authID").(uint64)

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
