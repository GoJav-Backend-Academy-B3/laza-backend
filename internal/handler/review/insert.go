package review

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
)

func (ct *reviewHandler) post(ctx *gin.Context) {
	//userID := ctx.get("authID")
	userId := uint64(3)
	productId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var review model.Review
	if err := json.NewDecoder(ctx.Request.Body).Decode(&review); err != nil {
		return
	}

	ct.insertReview.Execute(userId, productId, review.Comment, review.Rating).Send(ctx)
}
