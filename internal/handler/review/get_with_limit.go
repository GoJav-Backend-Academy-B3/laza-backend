package review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
	"github.com/phincon-backend/laza/helper"
)

// GetLimitReviewByIdProduct godoc
// @Summary Get bank identified by the given id
// @Description get the bank corresponding to the input Id
// @Tags ShowReviewLimit
// @Accept json
// @Produce json
// @Param id path int true "ID of the Product"
// @Security JWT
// @Success 200 {object} helper.Response{isError=bool,status=string,data=model.ProductReview}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /products/{id}/reviews/ [get]
func (h *reviewHandler) getWithLimit(c *gin.Context) {
	// productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	// if err != nil {
	// 	helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	// }
	// page := c.Query("page")
	// perpage := c.Query("perpage")

	// pageParse, _ := strconv.ParseUint(page, 10, 64)
	// perpageParse, _ := strconv.ParseUint(perpage, 10, 64)

	// h.getWithLimitReview.Execute(pageParse, perpageParse, productId).Send(c)
	productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	limit_q := c.DefaultQuery(QUERY_LIMIT, defaults.Query.LimitString())
	offset_q := c.DefaultQuery(QUERY_OFFSET, defaults.Query.OffsetString())

	limit, err := strconv.ParseUint(limit_q, 10, 32)
	if err != nil {
		limit = defaults.Query.Limit()
	}

	// convert offset to unsigned integer
	offset, err := strconv.ParseUint(offset_q, 10, 32)
	// if this fails, set to default value
	if err != nil {
		offset = defaults.Query.Offset()
	}
	h.getWithLimitReview.Execute(offset, limit, productId).Send(c)
}
