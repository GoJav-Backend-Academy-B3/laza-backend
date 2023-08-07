package review

import (
	"github.com/phincon-backend/laza/helper"
)

type InsertReviewUsecase interface {
	Execute(userId uint64, productId uint64, comment string, Rating float32) *helper.Response
}
