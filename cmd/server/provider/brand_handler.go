package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/brand"
	repository "github.com/phincon-backend/laza/internal/repo/brand"
	usecase "github.com/phincon-backend/laza/internal/usecase/brand"
)

func NewBrandHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	brandRepo := repository.NewBrandRepo(gorm)

	createBrandUsecase := usecase.NewCreateBrandUseCaseImpl(*brandRepo)
	getBrandByIdUsecase := usecase.NewGetBrandByIdUsecaseImpl(*brandRepo)
	searchBrandByName := usecase.NewSearchBrandByNameUsecaseImpl(*brandRepo)
	viewBrandUsecase := usecase.NewViewBrandUsecaseImpl(*brandRepo)
	deleteBrandUsecase := usecase.NewDeleteBrandUsecaseImpl(*brandRepo)
	updateBrandUsecase := usecase.NewUpdateBrandImpl(*brandRepo)

	validate := validator.New()

	return handler.NewBrandHandler(
		"/brand",
		createBrandUsecase,
		searchBrandByName,
		deleteBrandUsecase,
		getBrandByIdUsecase,
		updateBrandUsecase,
		viewBrandUsecase,
		validate,
	)
}
