package review

import (
	"math"

	review "github.com/phincon-backend/laza/domain/repositories/Review"
	"github.com/phincon-backend/laza/domain/response"
	action "github.com/phincon-backend/laza/domain/usecases/review"
	"github.com/phincon-backend/laza/helper"
)

type GetAllReviewUsecase struct {
	getReviewByProduct review.GetReviewByProduct
}

func NewGetAllReviewUsecase(repoProduct review.GetReviewByProduct) action.GetReviewByProductUsecase {
	return &GetAllReviewUsecase{
		getReviewByProduct: repoProduct,
	}
}

func (uc *GetAllReviewUsecase) Execute(productID uint64) *helper.Response {
	product, err := uc.getReviewByProduct.GetProductById(productID)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	averageRating, total, err := uc.getReviewByProduct.GetReviewStatsByProduct(productID)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}
	averageRating = math.Round(averageRating*100) / 100
	newProdustReviews := response.ProductReview{
		Total:       total,
		Products:    product,
		Avrg_Rating: averageRating,
	}

	return helper.GetResponse(newProdustReviews, 200, false)
}
