package wishlist

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

// Get Wishlist godoc
// @Summary Get Wishlist
// @Description Get Wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.WishListProductLimit}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /wishlists [GET]
func (h *getWishlistHandler) getByLimit(c *gin.Context) {

	userId := c.MustGet("userId").(uint64)

	// Get limit and offset query string
	limit_q := c.DefaultQuery(QUERY_LIMIT, defaults.Query.LimitString())
	offset_q := c.DefaultQuery(QUERY_OFFSET, defaults.Query.OffsetString())

	// convert limit to unsigned integer
	limit, err := strconv.ParseUint(limit_q, 10, 32)
	// if this fails, set to default value
	if err != nil {
		limit = defaults.Query.Limit()
	}

	// convert offset to unsigned integer
	offset, err := strconv.ParseUint(offset_q, 10, 32)
	// if this fails, set to default value
	if err != nil {
		offset = defaults.Query.Offset()
	}
	result_, err := h.getWishlistLimitUsecase.Execute(userId, offset, limit)
	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(c)
		return
	}

	var _data response.ProductOverview
	// response data
	var _response_data response.WishListProductLimit
	_response_data.Total = len(result_)
	for _, v := range result_ {

		// model product to response product
		_data.FillFromEntity(v)
		_response_data.Product = append(_response_data.Product, _data)
	}

	helper.GetResponse(_response_data, 200, false).Send(c)
}
