package provider

import (
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/review"
	repo "github.com/phincon-backend/laza/internal/repo/review"
	usecase "github.com/phincon-backend/laza/internal/usecase/review"
)

func NewReviewHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	reviewrepo := repo.NewReviewRepo(gorm)
	getAllReviewByProduct := usecase.NewGetAllReviewUsecase(reviewrepo)
	insertReview := usecase.NewinsertReviewUsecase(reviewrepo)
	getWithLimitReview := usecase.NewGetWithLimitReviewUsecase(reviewrepo)

	return handler.NewReviewHandler(getAllReviewByProduct, insertReview, getWithLimitReview)

}
