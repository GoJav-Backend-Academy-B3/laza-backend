package category

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

// UpdateCategoryNameById godoc
// @Summary Update the category's name based on the given id.
// @Description It updates the category's name. It updates the category that has the given ID.
// @Tags category
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=response.CategorySimpleResponse}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 404 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /category [put]
func (ch *categoryHandler) updateNameById(c *gin.Context) {
	var categoryRequest requests.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	result, err := ch.updateCategoryNameByIdUsecase.Execute(categoryRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.GetResponse(err.Error(), http.StatusNotFound, true).Send(c)
			return
		}
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	helper.GetResponse(result, http.StatusOK, false).Send(c)
	return
}
