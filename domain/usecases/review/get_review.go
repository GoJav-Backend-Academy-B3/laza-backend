package review

import (
	"github.com/phincon-backend/laza/helper"
)

type GetReviewByProductUsecase interface {
	Execute(productID uint64) *helper.Response
}
