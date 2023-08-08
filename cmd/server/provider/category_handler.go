package provider

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/category"
	"github.com/phincon-backend/laza/internal/repo/category"
	uc "github.com/phincon-backend/laza/internal/usecase/category"
)

func NewCategoryHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	categoryRepo := category.NewCategoryRepo(gorm)

	createCategoryUsecase := uc.NewCreateCategoryUsecaseImpl(*categoryRepo)
	deleteCategoryByIdUsecase := uc.NewDeleteCategoryByIdUsecaseImpl(*categoryRepo)
	getCategoryByIdUsecase := uc.NewGetCategoryByIdUsecaseImpl(categoryRepo)
	searchCategoryByNameUsecase := uc.NewSearchCategoryByNameUsecaseImpl(categoryRepo)
	updateCategoryNameByIdUsecase := uc.NewUpdateCategoryNameByIdUsecaseImpl(categoryRepo)
	viewCategoryUsecase := uc.NewViewCategoryUsecaseImpl(categoryRepo)

	return handler.NewCategoryHandler(
		"/category",
		createCategoryUsecase,
		deleteCategoryByIdUsecase,
		getCategoryByIdUsecase,
		searchCategoryByNameUsecase,
		updateCategoryNameByIdUsecase,
		viewCategoryUsecase,
	)
}
