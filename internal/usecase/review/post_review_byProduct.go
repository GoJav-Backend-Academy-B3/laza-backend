package review

import (
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	d "github.com/phincon-backend/laza/domain/repositories/cart"
	usecase "github.com/phincon-backend/laza/domain/usecases/review"
	h "github.com/phincon-backend/laza/helper"
)

type insertReviewUsecase struct {
	insertReviewRepo repositories.InsertAction[m.Review]
	getCartByIdRepo  d.GetCartByIdAction
}

func NewinsertReviewUsecase(review repositories.InsertAction[m.Review], gcr d.GetCartByIdAction) usecase.InsertReviewUsecase {
	return &insertReviewUsecase{
		insertReviewRepo: review,
		getCartByIdRepo:  gcr,
	}
}

func (uc *insertReviewUsecase) Execute(userId uint64, productId uint64, comment string, rating float32) *h.Response {
	rs, err := uc.getCartByIdRepo.GetCartById(userId)
	if len(rs) == 0 {
		return h.GetResponse("You must checkout before giving a review", http.StatusBadRequest, true)
	}
	if rating > 5 {
		return h.GetResponse("Rating cannot exceed 5", http.StatusBadRequest, true)
	}
	review := m.Review{
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
