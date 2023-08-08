package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// GetAllAddress godoc
// @Summary Get All Address
// @Description Get all Address
// @Tags address
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Address}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /address [get]
func (h *addressHandler) GetAllAddressByUserIdHandler(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	addresses, err := h.get.GetAllAddressByUserId(userId)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(addresses, http.StatusOK, false).Send(ctx)

}
