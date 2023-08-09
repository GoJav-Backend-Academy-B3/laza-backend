package review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// GetReviewByIdProduct godoc
// @Summary Get bank identified by the given id
// @Description get the bank corresponding to the input Id
// @Tags Review
// @Accept json
// @Produce json
// @Param id path int true "ID of the Product"
// @Security JWT
// @Success 200 {object} helper.Response{isError=bool,status=string,data=model.ProductReview}
// @Failure 404 {object} helper.Response{status=string,description=string,isError=bool}

// @Router /products/{id}/reviews/ [get]

func (h *reviewHandler) get(c *gin.Context) {
	productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	h.getAllReview.Execute(productId).Send(c)
}
