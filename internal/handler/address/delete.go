package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *addressHandler) DeleteAddressHandler(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := h.delete.DeleteAddressById(id)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse("Deleted Succesfully", http.StatusOK, false).Send(ctx)

}
