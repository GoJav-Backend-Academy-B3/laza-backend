package brand

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// DeleteBrand godoc
// @Summary Delete brand identified by the given id
// @Description Delete the brand corresponding to the input Id
// @Tags brand
// @Accept json
// @Produce json
// @Param id path int true "ID of the brand"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /brand/{id} [delete]
func (h *brandHandler) DeleteBrandById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	err = h.deleteBrandUsecase.Execute(uint64(id))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	helper.GetResponse("Delete Succesfully", http.StatusOK, true).Send(ctx)
}
