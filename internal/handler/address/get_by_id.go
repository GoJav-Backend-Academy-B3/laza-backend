package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// GetByIdAddress godoc
// @Summary Get Details for a given id
// @Description Get details of address corresponding is the input Id
// @Tags address
// @Accept json
// @Produce json
// @Param id path int true "ID of the address"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Address}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Error 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /address/{id} [get]
func (h *addressHandler) GetAddressByIdHandler(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	addresses, err := h.get.GetAddressById(id)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(addresses, http.StatusOK, false).Send(ctx)

}
