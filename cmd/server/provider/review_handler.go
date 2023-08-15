package provider

import (
	"github.com/go-playground/validator/v10"
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/review"
	repos "github.com/phincon-backend/laza/internal/repo/cart"
	repo "github.com/phincon-backend/laza/internal/repo/review"
	usecase "github.com/phincon-backend/laza/internal/usecase/review"
)

func NewReviewHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs
	validate := validator.New()
	reviewrepo := repo.NewReviewRepo(gorm)
	cartrepo := repos.NewCartRepo(gorm)
	getAllReviewByProduct := usecase.NewGetAllReviewUsecase(reviewrepo)
	insertReview := usecase.NewinsertReviewUsecase(reviewrepo, cartrepo)
	getWithLimitReview := usecase.NewGetWithLimitReviewUsecase(reviewrepo)

	return handler.NewReviewHandler(getAllReviewByProduct, insertReview, getWithLimitReview, validate)

}
