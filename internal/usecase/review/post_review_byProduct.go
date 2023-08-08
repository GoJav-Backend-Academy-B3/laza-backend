package review

import (
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	usecase "github.com/phincon-backend/laza/domain/usecases/review"
	h "github.com/phincon-backend/laza/helper"
)

type insertReviewUsecase struct {
	insertReviewRepo d.InsertAction[m.Review]
}

func NewinsertReviewUsecase(h d.InsertAction[m.Review]) usecase.InsertReviewUsecase {
	return &insertReviewUsecase{
		insertReviewRepo: h,
	}
}

func (uc *insertReviewUsecase) Execute(userId uint64, productId uint64, comment string, rating float32) *h.Response {
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

	return h.GetResponse(addreview, http.StatusOK, false)
}
