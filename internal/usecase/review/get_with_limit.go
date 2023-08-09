package review

import (
	review "github.com/phincon-backend/laza/domain/repositories/Review"
	rw "github.com/phincon-backend/laza/domain/usecases/review"
	"github.com/phincon-backend/laza/helper"
)

type GetWithLimitReviewUsecase struct {
	getWithLimitAction review.GetWithLimitReviewsAction
}

func NewGetWithLimitReviewUsecase(repo review.GetWithLimitReviewsAction) rw.GetWithLimitReviewUsecase {
	return &GetWithLimitReviewUsecase{getWithLimitAction: repo}
}

func (uc *GetWithLimitReviewUsecase) Execute(offset, limit uint64, productID uint64) *helper.Response {
	result, err := uc.getWithLimitAction.GetWithLimit(offset, limit, productID)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, false)
}
