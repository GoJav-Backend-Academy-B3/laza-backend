package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// PostCategory godoc
// @Summary Create new category
// @Description It creates a new Category record.
// @Tags category
// @Accept json
// @Produce json
// @Security JWT
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=response.CategorySimpleResponse}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /category [post]
func (ch *categoryHandler) postCategory(c *gin.Context) {
	var categoryRequest requests.CategoryRequest
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	result, err := ch.createCategoryUsecase.Execute(categoryRequest)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(result, http.StatusCreated, false).Send(c)
	return
}
