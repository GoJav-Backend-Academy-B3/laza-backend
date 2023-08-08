package review

import "github.com/phincon-backend/laza/helper"

type GetWithLimitReviewUsecase interface {
	Execute(limit, offset uint64, productID uint64) *helper.Response
}
