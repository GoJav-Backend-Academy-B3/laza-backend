package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
	"strconv"
)

// DeleteCategoryById godoc
// @Summary Delete category by id
// @Description Delete category by id
// @Tags category
// @Accept json
// @Param id path int true "Category ID"
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 404 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /category/:id [delete]
func (ch *categoryHandler) deleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}
	rowAffected, err := ch.deleteCategoryByIdUsecase.Execute(uint64(id))
	if err != nil || rowAffected == 0 {
		if rowAffected == 0 {
			helper.GetResponse("no record found", http.StatusNotFound, true).Send(c)
			return
		}
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(fmt.Sprintf("rows affected %d", rowAffected), http.StatusOK, false).Send(c)
	return

}
