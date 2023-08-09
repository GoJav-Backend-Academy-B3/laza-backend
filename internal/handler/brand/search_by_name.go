package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// DeleteAddress godoc
// @Summary Delete brand identified by the given id
// @Description Delete the brand corresponding to the input Id
// @Tags brand
// @Accept json
// @Produce json
// @Param name query string true "name of the brand"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Brand}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /brand/search [get]
func (h *brandHandler) SearchByBrandName(ctx *gin.Context) {
	name := ctx.Query("name")

	results, err := h.searchBrandByNameUsecase.Execute(name)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(results, http.StatusOK, false).Send(ctx)

}
