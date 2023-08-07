package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// DeleteAddress godoc
// @Summary Delete address identified by the given id
// @Description Delete the address corresponding to the input Id
// @Tags address
// @Accept json
// @Produce json
// @Param id path int true "ID of the address"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Address}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Error 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /address/{id} [delete]
func (h *addressHandler) DeleteAddressHandler(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := h.delete.DeleteAddressById(id)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse("Deleted Succesfully", http.StatusOK, false).Send(ctx)

}
