package wishlist

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
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
	h.getWishlistLimitUsecase.Execute(userId, offset, limit).Send(c)
}
