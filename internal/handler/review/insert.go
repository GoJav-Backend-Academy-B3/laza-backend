package review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

func (ct *reviewHandler) post(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	// userId := uint64(3)
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse("Invalid product ID", http.StatusBadRequest, true).Send(ctx)
		return
	}

	var review requests.ReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		helper.GetResponse("Invalid JSON data", http.StatusBadRequest, true).Send(ctx)
		return
	}

	// Validate the review model
	if err := ct.validate.Struct(review); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	ct.insertReview.Execute(userId, productId, review.Comment, review.Rating).Send(ctx)
}
