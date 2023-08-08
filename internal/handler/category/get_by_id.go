package category

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// Get Category By Id godoc
// @Summary Get category by id
// @Description Get category by its id
// @Tags category
// @Accept json
// @Produce json
// @Security JWT
// @Param id path int true "Category ID"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=response.CategorySimpleResponse}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 404 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /category/:id [get]
func (ch *categoryHandler) getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	result, err := ch.getCategoryByIdUsecase.Execute(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.GetResponse(err.Error(), http.StatusNotFound, true).Send(c)
			return
		}
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(result, http.StatusOK, false).Send(c)
	return
}
