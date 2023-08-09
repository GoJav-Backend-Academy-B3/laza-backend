package category

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

// GetAllCategory godoc
// @Summary Get All Category
// @Description Get all category
// @Tags category
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=[]response.CategorySimpleResponse}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /category/all [get]
func (ch *categoryHandler) getAll(c *gin.Context) {
	results, err := ch.viewCategoryUsecase.Execute()
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}
	helper.GetResponse(results, http.StatusOK, false).Send(c)
	return
}
