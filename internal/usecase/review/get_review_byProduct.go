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
	if len(product) == 0 {
		return helper.GetResponse("Reviews not found", 404, true)
	}
	averageRating, total, err := uc.getReviewByProduct.GetReviewStatsByProduct(productID)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}
	averageRating = math.Round(averageRating*10) / 10
	newProdustReviews := response.GetReviews{
		Total:       total,
		Reviews:     product,
		Avrg_Rating: averageRating,
	}

	return helper.GetResponse(newProdustReviews, 200, false)
}
