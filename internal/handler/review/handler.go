package review

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/review"
	"github.com/phincon-backend/laza/middleware"
)

type reviewHandler struct {
	getAllReview       review.GetReviewByProductUsecase
	insertReview       review.InsertReviewUsecase
	getWithLimitReview review.GetWithLimitReviewUsecase
	validate           *validator.Validate
}

func NewReviewHandler(
	getAllReview review.GetReviewByProductUsecase,
	insertReview review.InsertReviewUsecase,
	getWithLimitReview review.GetWithLimitReviewUsecase,
	validate *validator.Validate,

) handlers.HandlerInterface {
	return &reviewHandler{
		getAllReview:       getAllReview,
		insertReview:       insertReview,
		getWithLimitReview: getWithLimitReview,
		validate:           validate,
	}
}

func (h *reviewHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/products/:id/reviews",
			HandlerFunc: h.get},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/products/:id/reviews",
			HandlerFunc: h.post,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/products/:id/review",
			HandlerFunc: h.getWithLimit},
	)

	return
}
