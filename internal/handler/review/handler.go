package review

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/review"
)

type reviewHandler struct {
	getAllReview       review.GetReviewByProductUsecase
	insertReview       review.InsertReviewUsecase
	getWithLimitReview review.GetWithLimitReviewUsecase
}

func NewReviewHandler(
	getAllReview review.GetReviewByProductUsecase,
	insertReview review.InsertReviewUsecase,
	getWithLimitReview review.GetWithLimitReviewUsecase,

) handlers.HandlerInterface {
	return &reviewHandler{
		getAllReview:       getAllReview,
		insertReview:       insertReview,
		getWithLimitReview: getWithLimitReview,
	}
}

func (h *reviewHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/products/:id/reviews", HandlerFunc: h.get},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/products/:id/reviews", HandlerFunc: h.post},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/products/:id/reviews/", HandlerFunc: h.getWithLimit},
	)

	return
}
