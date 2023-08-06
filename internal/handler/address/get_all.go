package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *addressHandler) GetAllAddressByUserIdHandler(ctx *gin.Context) {
	userId, _ := strconv.ParseUint(ctx.Param("userId"), 10, 32)

	addresses, err := h.get.GetAllAddressByUserId(userId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(addresses, http.StatusOK, false).Send(ctx)

}
