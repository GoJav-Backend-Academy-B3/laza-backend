package category

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

// FindCategoryByName godoc
// @Summary Find category by its name
// @Description It returns a list of categories. This is not an exact match searching method.
// @Tags category
// @Accept json
// @Produce json
// @Security JWT
// @Param name query string true "Category name"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=[]response.CategorySimpleResponse}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /category/search [get]
func (ch *categoryHandler) searchByName(c *gin.Context) {
	name := c.Query("name")
	results, err := ch.searchCategoryByNameUsecase.Execute(name)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}
	helper.GetResponse(results, http.StatusOK, false).Send(c)
	return
}
