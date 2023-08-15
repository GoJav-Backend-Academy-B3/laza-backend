package review

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	cart "github.com/phincon-backend/laza/domain/repositories/cart"
	usecase "github.com/phincon-backend/laza/domain/usecases/review"
	h "github.com/phincon-backend/laza/helper"
)

type insertReviewUsecase struct {
	insertReviewRepo repositories.InsertAction[model.Review]
	getCartByIdRepo  cart.GetCartByIdAction
}

func NewinsertReviewUsecase(reviews repositories.InsertAction[model.Review], carts cart.GetCartByIdAction) usecase.InsertReviewUsecase {
	return &insertReviewUsecase{
		insertReviewRepo: reviews,
		getCartByIdRepo:  carts,
	}
}

func (uc *insertReviewUsecase) Execute(userId uint64, productId uint64, comment string, rating float32) *h.Response {
	if rating > 5 {
		return h.GetResponse("Rating cannot exceed 5", http.StatusBadRequest, true)
	}
	review := model.Review{
		UserId:    userId,
		ProductId: productId,
		Comment:   comment,
		Rating:    rating,
	}

	addreview, err := uc.insertReviewRepo.Insert(review)
	if err != nil {
		return h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return h.GetResponse(addreview, 201, false)
}
