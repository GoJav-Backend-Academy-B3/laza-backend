package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// CreateAddress godoc
// @Summary Post Details
// @Description Post details of address
// @Tags address
// @Accept json
// @Produce json
// @Param address body requests.AddressRequest true "create address"
// @Security JWT
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=model.Address}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Error 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /address [post]
func (h *addressHandler) PostAddressHandler(ctx *gin.Context) {
	var request requests.AddressRequest

	userId := ctx.MustGet("userId").(uint64)

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
