package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// GetAllBrand godoc
// @Summary Get All brand
// @Description Get all brand
// @Tags brand
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Brand}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /brand [get]
func (h *brandHandler) ViewAllBrand(ctx *gin.Context) {
	result, err := h.viewBrandUsecase.Execute()
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(result, http.StatusOK, true).Send(ctx)

}
